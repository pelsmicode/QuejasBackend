package api

import (
	"github.com/gorilla/mux"
	"github.com/pelsmicode/QuejasBackend/db"
	"github.com/pelsmicode/QuejasBackend/handler"
	"github.com/pelsmicode/QuejasBackend/repository"
	"github.com/pelsmicode/QuejasBackend/service"
)

var (
	client        = db.NewSqlClient()
	regionHandler = handler.RegionHandler{S: service.NewRegionService(repository.NewRegionRepository(client.DB))}
)

func Controllers() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/region", regionHandler.GetAllRegions).Methods("GET")

	return router
}
