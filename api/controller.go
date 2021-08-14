package api

import (
	"github.com/gorilla/mux"
	"github.com/pelsmicode/QuejasBackend/db"
	"github.com/pelsmicode/QuejasBackend/handler"
	"github.com/pelsmicode/QuejasBackend/repository"
	"github.com/pelsmicode/QuejasBackend/service"
)

var (
	client           = db.NewSqlClient()
	regionHandler    = handler.RegionHandler{S: service.NewRegionService(repository.NewRegionRepository(client.DB))}
	deparmentHandler = handler.DeparmentHandler{S: service.NewDeparmentService(repository.NewDepartmentRepository(client.DB))}
	townshipHandler  = handler.TownshipHandler{S: service.NewTownshipService(repository.NewTownshipRepository(client.DB))}
	branchHandler    = handler.BranchHandler{S: service.NewBranchService(repository.NewBranchRepository(client.DB))}
)

func Controllers() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/region", regionHandler.GetAllRegions).Methods("GET")
	router.HandleFunc("/region/{id}", regionHandler.GetRegion).Methods("GET")
	router.HandleFunc("/deparment", deparmentHandler.GetAllDepartments).Methods("GET")
	router.HandleFunc("/deparment/{id:[0-9]+}", deparmentHandler.GetDeparment).Methods("GET")
	router.HandleFunc("/township", townshipHandler.GetAllTownships).Methods("GET")
	router.HandleFunc("/township/{id:[0-9]+}", townshipHandler.GetTownshipByID).Methods("GET")
	router.HandleFunc("/township/deparment/{id:[0-9]+}", townshipHandler.GetTownshipsByDep).Methods("GET")
	router.HandleFunc("/diaco", branchHandler.GetDiacoBranches).Methods("GET")

	return router
}
