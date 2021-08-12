package handler

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pelsmicode/QuejasBackend/service"
)

type DeparmentHandler struct {
	S service.DepartmentService
}

func (h *DeparmentHandler) GetAllDepartments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	deparments, err := h.S.GetDeparments()
	if err != nil {
		log.Println("[Handler Error GetDeparments]", err)
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	writeResponse(w, http.StatusOK, deparments)
}

func (h *DeparmentHandler) GetDeparment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("[Handler Error GetDeparmentByID]", err)
		writeResponse(w, http.StatusNotFound, errors.New("Error in ID").Error())
		return
	}

	deparment, err := h.S.GetDeparmentByID(id)
	if err != nil {
		log.Println("[Handler Error GetDeparmentByID]", err)
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	writeResponse(w, http.StatusOK, deparment)
}
