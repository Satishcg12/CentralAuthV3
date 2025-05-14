package middlewares

import (
	"net/http"

	"github.com/Satishcg12/CentralAuthV3/server/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// ValidationMiddleware returns a middleware that handles validation errors
func (m *Middleware) ValidationMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Set the validator to the context if not already set
			if c.Echo().Validator == nil {
				c.Echo().Validator = utils.NewValidator()
			}
			// Continue with the next middleware or handler
			if err := next(c); err != nil {
				// If it's a validation error (HTTP error with validation info)
				if he, ok := err.(*echo.HTTPError); ok && he.Code == http.StatusBadRequest {
					// Check if the error is from the validator
					if validationErrs, ok := he.Internal.(validator.ValidationErrors); ok {
						// Format the validation errors using the utility function
						validationErrors := utils.FormatValidationErrors(validationErrs)

						// Return a standardized validation error response
						return utils.RespondWithError(
							c,
							utils.StatusCodeBadRequest,
							"Validation failed",
							utils.ErrorCodeValidationFailed,
							"Input validation failed",
							validationErrors,
						)
					}
				}

				// For other types of errors, pass them through
				return err
			}

			return nil
		}
	}
}
