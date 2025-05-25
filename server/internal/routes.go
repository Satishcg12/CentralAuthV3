package internal

import (
	"github.com/Satishcg12/CentralAuthV3/server/internal/config"
	"github.com/Satishcg12/CentralAuthV3/server/internal/db"
	"github.com/Satishcg12/CentralAuthV3/server/internal/features"
	"github.com/Satishcg12/CentralAuthV3/server/internal/features/auth"
	"github.com/Satishcg12/CentralAuthV3/server/internal/features/client"
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
	clientHandler := client.NewClientHandler(ah)

	// API v1 group - Register API routes FIRST
	v1 := e.Group("/api/v1")

	// Health check - public
	v1.GET("/health", healthHandler.Check)

	// Auth Endpoints - Public
	v1.POST("/auth/register", authHandler.Register) // User registration

	// Client Endpoints - TODO: Add authentication middleware when available
	v1.POST("/clients", clientHandler.CreateClient)                                                               // Create new client
	v1.GET("/clients", clientHandler.GetAllClients)                                                               // List all clients
	v1.GET("/clients/:id", clientHandler.GetClientById)                                                           // Get client by UUID
	v1.PUT("/clients/:id", clientHandler.UpdateClient)                                                            // Update client by UUID
	v1.DELETE("/clients/:id", clientHandler.DeleteClient)                                                         // Delete client by UUID
	v1.POST("/clients/:id/regenerate-secret", clientHandler.RegenerateClientSecret)                               // Regenerate secret by UUID
	v1.POST("/clients/by-client-id/:client_id/regenerate-secret", clientHandler.RegenerateClientSecretByClientID) // Regenerate secret by client_id

	// Static file serving for assets - MUST come before SPA fallback
	e.Static("/assets", "./dist/assets")

	// Serve specific static files from dist root
	e.File("/favicon.ico", "./dist/favicon.ico")
	e.File("/manifest.json", "./dist/manifest.json")
	e.File("/robots.txt", "./dist/robots.txt")
	e.File("/logo.png", "./dist/logo.png")
	e.File("/logo192.png", "./dist/logo192.png")
	e.File("/logo2.png", "./dist/logo2.png")
	e.File("/logo512.png", "./dist/logo512.png")

	// SPA fallback - serve index.html for all other routes that are not API or assets
	e.GET("/*", func(c echo.Context) error {
		path := c.Request().URL.Path

		// Don't serve index.html for API routes
		if len(path) >= 4 && path[:4] == "/api" {
			return echo.NewHTTPError(404, "API endpoint not found")
		}

		// Don't serve index.html for assets - let them 404 properly
		if len(path) >= 7 && path[:7] == "/assets" {
			return echo.NewHTTPError(404, "Asset not found")
		}

		// Serve index.html for all other routes (SPA routing)
		return c.File("./dist/index.html")
	})

}
