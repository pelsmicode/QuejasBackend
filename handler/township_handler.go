package handler

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pelsmicode/QuejasBackend/service"
)

type TownshipHandler struct {
	S service.TownshipService
}

func (h *TownshipHandler) GetAllTownships(w http.ResponseWriter, r *http.Request) {
	township, err := h.S.GetTownships()
	if err != nil {
		log.Println("[Handler Error GetTownships]", err)
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	writeResponse(w, http.StatusOK, township)
}

func (h *TownshipHandler) GetTownshipByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("[Handler Error GetTownshipByID]", err)
		writeResponse(w, http.StatusNotFound, errors.New("Error in ID").Error())
		return
	}

	township, err := h.S.GetTownshipByID(id)
	if err != nil {
		log.Println("[Handler Error GetTownshipByID]", err)
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	writeResponse(w, http.StatusAccepted, township)
}
