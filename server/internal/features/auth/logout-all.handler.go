package auth

import (
	"net/http"

	"github.com/Satishcg12/CentralAuthV3/server/internal/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *AuthHandler) LogoutAll(c echo.Context) error {
	// Parse the request body (though it might be empty)
	req := new(LogoutAllRequest)
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

	// Get user ID from context (set by auth middleware)
	userIDStr, ok := c.Get("user_id").(string)
	if !ok {
		return utils.RespondWithError(
			c,
			utils.StatusCodeUnauthorized,
			"Unauthorized",
			utils.ErrorCodeUnauthorized,
			"User not authenticated",
			nil,
		)
	}

	// Parse UUID
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return utils.RespondWithError(
			c,
			utils.StatusCodeBadRequest,
			"Invalid user ID",
			utils.ErrorCodeInvalidRequest,
			"Invalid user ID format",
			err,
		)
	}

	// Count active sessions before deactivating them
	activeSessions, err := h.store.CountActiveUserSessions(c.Request().Context(), userID)
	if err != nil {
		return utils.RespondWithError(
			c,
			utils.StatusCodeInternalError,
			"Internal Server Error",
			utils.ErrorCodeDatabaseError,
			"Failed to count active sessions",
			err,
		)
	}

	// Deactivate all sessions for the user
	err = h.store.DeactivateAllUserSessions(c.Request().Context(), userID)
	if err != nil {
		return utils.RespondWithError(
			c,
			utils.StatusCodeInternalError,
			"Internal Server Error",
			utils.ErrorCodeDatabaseError,
			"Failed to deactivate sessions",
			err,
		)
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
	res := LogoutAllResponse{
		Success:       true,
		Message:       "All sessions terminated successfully",
		SessionsEnded: int(activeSessions),
	}

	return utils.RespondWithSuccess(
		c,
		utils.StatusCodeSuccess,
		"All sessions terminated successfully",
		res,
	)
}
