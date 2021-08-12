package api

import (
	"github.com/pelsmicode/QuejasBackend/db"
	"github.com/pelsmicode/QuejasBackend/handler"
	"github.com/pelsmicode/QuejasBackend/repository"
	"github.com/pelsmicode/QuejasBackend/service"
)

var (
	client        = db.NewSqlClient()
	regionHandler = handler.RegionHandler{S: service.NewRegionService(repository.NewRegionRepository(client.DB))}
)
