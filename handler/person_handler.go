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

	if p.Township == 0 {
		log.Println("[Handler] Error SavePerson", "No Twonship ID")
		writeResponse(w, http.StatusBadRequest, p.Township)
		return
	}

	if p.Branch == 0 {
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
