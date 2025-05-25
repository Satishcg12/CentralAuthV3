package client

import (
	"github.com/Satishcg12/CentralAuthV3/server/internal/config"
	"github.com/Satishcg12/CentralAuthV3/server/internal/db"
	"github.com/Satishcg12/CentralAuthV3/server/internal/features"
)

type ClientHandler struct {
	store  *db.Store
	config *config.Config
	DB     *db.Store // Add this for compatibility with existing handlers
}

// NewClientHandler creates a new client handler
func NewClientHandler(ah *features.AppHandlers) *ClientHandler {
	return &ClientHandler{
		store:  ah.Store,
		config: ah.Cfg,
		DB:     ah.Store, // Set both for compatibility
	}
}
