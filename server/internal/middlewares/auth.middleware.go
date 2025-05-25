package middlewares

import (
	"strings"

	"github.com/Satishcg12/CentralAuthV3/server/internal/utils"
	"github.com/labstack/echo/v4"
)

// AuthMiddleware creates a middleware that validates JWT tokens and extracts user information
func (m *Middleware) AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get token from Authorization header
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				// Try to get token from cookie as fallback
				cookie, err := c.Cookie("access_token")
				if err != nil {
					return utils.RespondWithError(
						c,
						utils.StatusCodeUnauthorized,
						"Unauthorized",
						utils.ErrorCodeUnauthorized,
						"Missing authorization token",
						err,
					)
				}
				authHeader = "Bearer " + cookie.Value
			}

			// Check if the header starts with "Bearer "
			if !strings.HasPrefix(authHeader, "Bearer ") {
				return utils.RespondWithError(
					c,
					utils.StatusCodeUnauthorized,
					"Unauthorized",
					utils.ErrorCodeUnauthorized,
					"Invalid authorization header format",
					nil,
				)
			}

			// Extract the token
			token := strings.TrimPrefix(authHeader, "Bearer ")

			// Validate the token
			claims, err := utils.ValidateToken(token)
			if err != nil {
				return utils.RespondWithError(
					c,
					utils.StatusCodeUnauthorized,
					"Unauthorized",
					utils.ErrorCodeUnauthorized,
					"Invalid or expired token",
					err,
				)
			}

			// Store user information in context
			c.Set("user_id", claims.UserID)
			c.Set("user_email", claims.Email)
			c.Set("user_claims", claims)

			return next(c)
		}
	}
}

// OptionalAuthMiddleware is similar to AuthMiddleware but doesn't fail if no token is provided
func (m *Middleware) OptionalAuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get token from Authorization header
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				// Try to get token from cookie as fallback
				cookie, err := c.Cookie("access_token")
				if err != nil {
					// No token found, continue without authentication
					return next(c)
				}
				authHeader = "Bearer " + cookie.Value
			}

			// Check if the header starts with "Bearer "
			if !strings.HasPrefix(authHeader, "Bearer ") {
				// Invalid format, continue without authentication
				return next(c)
			}

			// Extract the token
			token := strings.TrimPrefix(authHeader, "Bearer ")

			// Validate the token
			claims, err := utils.ValidateToken(token)
			if err != nil {
				// Invalid token, continue without authentication
				return next(c)
			}

			// Store user information in context
			c.Set("user_id", claims.UserID)
			c.Set("user_email", claims.Email)
			c.Set("user_claims", claims)

			return next(c)
		}
	}
}
