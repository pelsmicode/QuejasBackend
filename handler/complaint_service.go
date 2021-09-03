package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pelsmicode/QuejasBackend/model"
	"github.com/pelsmicode/QuejasBackend/service"
)

type ComplaintHandler struct {
	S service.ComplaintService
}

func (h *ComplaintHandler) SaveComplaint(w http.ResponseWriter, r *http.Request) {
	var c model.ComplaintRequest

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		log.Println("[Handler Error SaveComplaint]", err)
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.S.SaveComplaint(c)
	if err != nil {
		log.Println("[Handler Error SaveComplaint]", err)
		writeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeResponse(w, http.StatusCreated, c)
}

func (h *ComplaintHandler) GetMainComplaints(w http.ResponseWriter, r *http.Request) {
	c, err := h.S.GetMainComplaints()
	if err != nil {
		log.Println("[Handler Error GetMainComplaints]", err)
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	writeResponse(w, http.StatusOK, c)
}
