package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pelsmicode/QuejasBackend/service"
)

type RegionHandler struct {
	S service.RegionService
}

func (h *RegionHandler) GetAllRegions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	region, err := h.S.GetRegions()
	if err != nil {
		log.Println("[Handler Error GetRegions]", err)
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	writeResponse(w, http.StatusOK, region)
}

func (h *RegionHandler) GetRegion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("[Handler Error GetRegionByID]", err)
		writeResponse(w, http.StatusNotFound, errors.New("Error in ID").Error())
		return
	}

	region, err := h.S.GetRegionByID(id)
	if err != nil {
		log.Println("[Handler Error GetRegionByID]", err)
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	writeResponse(w, http.StatusAccepted, region)
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
