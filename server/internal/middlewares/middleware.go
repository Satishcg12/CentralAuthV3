package middlewares

import (
	"github.com/Satishcg12/CentralAuthV3/server/internal/config"
	"github.com/Satishcg12/CentralAuthV3/server/internal/db"
	"github.com/labstack/echo/v4"
)

// IMiddleware defines the interface for middleware operations
type IMiddleware interface {
	// Add middleware methods here
	ValidationMiddleware() echo.MiddlewareFunc
}

type Middleware struct {
	Store  *db.Store
	Config *config.Config
}

func NewMiddleware(m Middleware) IMiddleware {
	return &Middleware{
		Store:  m.Store,
		Config: m.Config,
	}
}

// Ensure Middleware implements IMiddleware
var _ IMiddleware = (*Middleware)(nil)
