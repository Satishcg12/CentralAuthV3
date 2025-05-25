package auth

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/Satishcg12/CentralAuthV3/server/internal/utils"
	"github.com/labstack/echo/v4"
)

func (h *AuthHandler) RefreshToken(c echo.Context) error {
	// Parse the request body (though it might be empty)
	req := new(RefreshTokenRequest)
	if err := c.Bind(req); err != nil {
		return utils.RespondWithError(
			c,
			utils.StatusCodeBadRequest,
			"Invalid request data",
			utils.ErrorCodeInvalidRequest,
			"Could not parse request body",
			err,
		)
	}

	// Get session token from cookie
	sessionCookie, err := c.Cookie("session_token")
	if err != nil {
		return utils.RespondWithError(
			c,
			utils.StatusCodeUnauthorized,
			"Unauthorized",
			utils.ErrorCodeUnauthorized,
			"No active session found",
			err,
		)
	}

	// Get the session from database to verify it exists and is active
	session, err := h.store.GetSessionByToken(c.Request().Context(), sessionCookie.Value)
	if err != nil {
		if err == sql.ErrNoRows {
			return utils.RespondWithError(
				c,
				utils.StatusCodeUnauthorized,
				"Unauthorized",
				utils.ErrorCodeUnauthorized,
				"Invalid or expired session",
				err,
			)
		}
		return utils.RespondWithError(
			c,
			utils.StatusCodeInternalError,
			"Internal Server Error",
			utils.ErrorCodeDatabaseError,
			"Failed to retrieve session",
			err,
		)
	}

	// Get user information from the session
	user, err := h.store.GetUserByID(c.Request().Context(), session.UserID)
	if err != nil {
		return utils.RespondWithError(
			c,
			utils.StatusCodeInternalError,
			"Internal Server Error",
			utils.ErrorCodeDatabaseError,
			"Failed to retrieve user",
			err,
		)
	}

	// Create new access token claims
	claims := utils.AccessTokenClaims{
		UserID:        user.ID.String(),
		Email:         user.Email,
		FullName:      user.FullName,
		EmailVerified: user.EmailVerified.Valid && user.EmailVerified.Bool,
	}

	// Create the new access token
	accessToken, expiresIn, err := utils.CreateAccessToken(claims)
	if err != nil {
		return utils.RespondWithError(
			c,
			utils.StatusCodeInternalError,
			"Internal Server Error",
			utils.ErrorCodeInternalError,
			"Failed to create access token",
			err,
		)
	}

	// Set the new access token in cookie
	c.SetCookie(&http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	// Calculate expiration timestamp
	expiresAt := time.Now().Add(time.Second * time.Duration(expiresIn)).Unix()

	// Create the response
	res := RefreshTokenResponse{
		AccessToken: accessToken,
		ExpiresAt:   expiresAt,
	}

	return utils.RespondWithSuccess(
		c,
		utils.StatusCodeSuccess,
		"Token refreshed successfully",
		res,
	)
}
