package auth

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/Satishcg12/CentralAuthV3/server/internal/db/sqlc"
	"github.com/Satishcg12/CentralAuthV3/server/internal/utils"
	"github.com/labstack/echo/v4"
)

func (h *AuthHandler) Login(c echo.Context) error {
	// Parse the request body
	req := new(LoginRequest)
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
	// Validate the request data
	if err := c.Validate(req); err != nil {
		return err
	}

	// check if user exists
	user, err := h.store.GetUserByEmail(c.Request().Context(), req.Email)
	if err != nil {
		return utils.RespondWithError(
			c,
			utils.StatusCodeUnauthorized,
			"Unauthorized",
			utils.ErrorCodeUnauthorized,
			"Invalid email or password",
			err,
		)
	}
	// Verify password
	if isValid := utils.ComparePasswords(user.PasswordHash, req.Password); !isValid {
		// If password is invalid, return unauthorized error
		return utils.RespondWithError(
			c,
			utils.StatusCodeUnauthorized,
			"Unauthorized",
			utils.ErrorCodeUnauthorized,
			"Invalid email or password",
			nil,
		)
	}
	// Create access token claims
	claims := utils.AccessTokenClaims{
		UserID:        user.ID.String(),
		Email:         user.Email,
		FullName:      user.FullName,
		EmailVerified: user.EmailVerified.Valid && user.EmailVerified.Bool,
	}

	// Create the access token
	accessToken, _, err := utils.CreateAccessToken(claims)
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

	sessionToken, err := utils.GenerateRandomString(32) // Generate a random session
	if err != nil {
		return utils.RespondWithError(
			c,
			utils.StatusCodeInternalError,
			"Internal Server Error",
			utils.ErrorCodeInternalError,
			"Failed to generate session token",
			err,
		)
	}
	userAgent := c.Request().UserAgent()
	ipAddress := c.RealIP()

	// create session in database
	session, err := h.store.CreateSession(c.Request().Context(), sqlc.CreateSessionParams{
		UserID:       user.ID,
		SessionToken: sessionToken,
		UserAgent:    sql.NullString{String: userAgent, Valid: userAgent != ""},
		IpAddress:    sql.NullString{String: ipAddress, Valid: ipAddress != ""},
		ExpiresAt:    time.Now().Add(30 * 24 * time.Hour), // set expired after a months
	})
	if err != nil {
		return utils.RespondWithError(
			c,
			utils.StatusCodeInternalError,
			"Internal Server Error",
			utils.ErrorCodeDatabaseError,
			"Failed to create session",
			err,
		)
	}

	// Set the session token in the cookie
	c.SetCookie(&http.Cookie{
		Name:     "session_token",
		Value:    session.SessionToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	// set access token in the cookie
	c.SetCookie(&http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	// Create the response
	res := LoginResponse{
		AccessToken: accessToken,
	}

	return utils.RespondWithSuccess(
		c,
		utils.StatusCodeSuccess,
		"Login successful",
		res,
	)

}
