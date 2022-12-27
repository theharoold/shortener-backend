package api

import (
	"github.com/theharoold/shortener-backend/config"
	"github.com/theharoold/shortener-backend/storage"
)

type API struct {
	Stg map[string]storage.Storage
	Cfg config.Config
}
