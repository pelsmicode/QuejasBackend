package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pelsmicode/QuejasBackend/service"
)

type RegionHandler struct {
	S service.RegionService
}

func (h *RegionHandler) GetAllRegions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	region, err := h.S.GetRegions()
	if err != nil {
		log.Println("[Handler Error GetRegion]", err)
		writeResponse(w, http.StatusBadRequest, err.Error())
	}

	writeResponse(w, http.StatusOK, region)
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
