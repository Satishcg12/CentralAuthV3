package internal

import (
	"github.com/Satishcg12/CentralAuthV3/server/internal/config"
	"github.com/Satishcg12/CentralAuthV3/server/internal/db"
	"github.com/Satishcg12/CentralAuthV3/server/internal/features"
	"github.com/Satishcg12/CentralAuthV3/server/internal/features/auth"
	"github.com/Satishcg12/CentralAuthV3/server/internal/features/health"
	"github.com/Satishcg12/CentralAuthV3/server/internal/middlewares"
	"github.com/labstack/echo/v4"
)

// setupRoutes configures all routes for the application
func SetupRoutes(e *echo.Echo, store *db.Store, cfg *config.Config, cm middlewares.IMiddleware) {
	ah := &features.AppHandlers{
		Store: store,
		Cfg:   cfg,
	}

	healthHandler := health.NewHealthHandler(ah)
	authHandler := auth.NewAuthHandler(ah)

	// API v1 group
	v1 := e.Group("/api/v1")

	// Health check - public
	v1.GET("/health", healthHandler.Check)

	// Auth Endpoints - Public
	v1.POST("/auth/register", authHandler.Register) // User registration

}
