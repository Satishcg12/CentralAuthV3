package features

import (
	"github.com/Satishcg12/CentralAuthV3/server/internal/config"
	"github.com/Satishcg12/CentralAuthV3/server/internal/db"
)

type AppHandlers struct {
	Store *db.Store
	Cfg   *config.Config
}
