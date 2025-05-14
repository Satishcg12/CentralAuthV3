package auth

import (
	"github.com/Satishcg12/CentralAuthV3/server/internal/config"
	"github.com/Satishcg12/CentralAuthV3/server/internal/db"
	"github.com/Satishcg12/CentralAuthV3/server/internal/features"
)

type AuthHandler struct {
	store  *db.Store
	config *config.Config
}

// NewAuthHandler creates a new authentication handler
func NewAuthHandler(ah *features.AppHandlers) *AuthHandler {
	return &AuthHandler{
		store:  ah.Store,
		config: ah.Cfg,
	}
}
