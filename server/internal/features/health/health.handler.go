package health

import (
	"github.com/Satishcg12/CentralAuthV3/server/internal/features"
	"github.com/Satishcg12/CentralAuthV3/server/internal/utils"
	"github.com/labstack/echo/v4"
)

// HealthHandler represents the health check handler
type HealthHandler struct{}

// NewHealthHandler creates a new health check handler
func NewHealthHandler(ah *features.AppHandlers) *HealthHandler {
	return &HealthHandler{}
}

// Check handles the health check endpoint
func (h *HealthHandler) Check(c echo.Context) error {
	return utils.RespondWithSuccess(
		c,
		utils.StatusCodeSuccess,
		"API is healthy",
		map[string]string{
			"status":  "up",
			"version": "1.0.0",
		},
	)
}
