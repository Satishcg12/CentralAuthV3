package internal

import (
	"net/http"

	"github.com/Satishcg12/CentralAuthV3/server/internal/config"
	"github.com/Satishcg12/CentralAuthV3/server/internal/middlewares"
	"github.com/labstack/echo/v4"
	emiddleware "github.com/labstack/echo/v4/middleware"
)

func SetupGlobalMiddleware(e *echo.Echo, cfg *config.Config, cm middlewares.IMiddleware) {

	// Add CORS middleware
	e.Use(emiddleware.CORSWithConfig(emiddleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:5173",
			"http://localhost:3000",
			"http://localhost:3001",
			"http://localhost:3002",
			"http://sso.local:3001",
			"http://myapp.local:3002",
		}, // Replace with your frontend URL
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
			echo.HeaderAccessControlAllowCredentials,
		},
		AllowCredentials: true,
		MaxAge:           86400, // 24 hours
	}))

	// Add other middleware
	e.Use(emiddleware.LoggerWithConfig(emiddleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status} \nmessage=${error}\n",
	}))
	e.Use(emiddleware.Recover())

	// Add validation middleware
	e.Use(cm.ValidationMiddleware())

}
