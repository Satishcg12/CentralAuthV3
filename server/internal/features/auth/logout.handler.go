package auth

import (
	"database/sql"
	"net/http"

	"github.com/Satishcg12/CentralAuthV3/server/internal/utils"
	"github.com/labstack/echo/v4"
)

func (h *AuthHandler) Logout(c echo.Context) error {
	// Parse the request body (though it might be empty)
	req := new(LogoutRequest)
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
			// Session doesn't exist or is already inactive
			// Still proceed with clearing cookies for client-side cleanup
		} else {
			return utils.RespondWithError(
				c,
				utils.StatusCodeInternalError,
				"Internal Server Error",
				utils.ErrorCodeDatabaseError,
				"Failed to retrieve session",
				err,
			)
		}
	} else {
		// Deactivate the session in database
		err = h.store.DeactivateSession(c.Request().Context(), session.SessionToken)
		if err != nil {
			return utils.RespondWithError(
				c,
				utils.StatusCodeInternalError,
				"Internal Server Error",
				utils.ErrorCodeDatabaseError,
				"Failed to deactivate session",
				err,
			)
		}
	}

	// Clear the session token cookie
	c.SetCookie(&http.Cookie{
		Name:     "session_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   -1, // This deletes the cookie
	})

	// Clear the access token cookie
	c.SetCookie(&http.Cookie{
		Name:     "access_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   -1, // This deletes the cookie
	})

	// Create the response
	res := LogoutResponse{
		Success: true,
		Message: "Logout successful",
	}

	return utils.RespondWithSuccess(
		c,
		utils.StatusCodeSuccess,
		"Logout successful",
		res,
	)
}
