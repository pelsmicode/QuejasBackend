package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pelsmicode/QuejasBackend/model"
	"github.com/pelsmicode/QuejasBackend/service"
)

type PersonHandler struct {
	S service.PersonService
}

func (h *PersonHandler) SavePerson(w http.ResponseWriter, r *http.Request) {
	var p model.PersonRequest

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Println("[Handler Error SavePerson]", err)
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if p.Township == "" {
		log.Println("[Handler] Error SavePerson", "No Twonship ID")
		writeResponse(w, http.StatusBadRequest, p.Township)
		return
	}

	if p.Branch == "" {
		log.Println("[Handler] Error SavePerson", "No Branch ID")
		writeResponse(w, http.StatusBadRequest, p.Branch)
		return
	}

	err = h.S.SavePerson(p)
	if err != nil {
		log.Println("[Handler Error SavePerson]", err)
		writeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeResponse(w, http.StatusCreated, model.ToPresonResponse(p))
}

func (h *PersonHandler) GetLastPerson(w http.ResponseWriter, r *http.Request) {
	var p model.PersonRequest
	id := h.S.GetLastPersonID()
	if id == 0 {
		log.Println("[Handler Error GetLastPerson]")
		writeResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	p.ID = id
	writeResponse(w, http.StatusOK, model.ToPresonResponse(p))
}
