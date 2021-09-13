package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pelsmicode/QuejasBackend/model"
	"github.com/pelsmicode/QuejasBackend/service"
)

type CompanyHandler struct {
	S service.CompanyService
}

func (h *CompanyHandler) SaveCompany(w http.ResponseWriter, r *http.Request) {
	var c model.CompanyRequest

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		log.Println("[Handler Error SaveCompany]", err)
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if c.Township == "" {
		log.Println("[Handler] Error SaveCompany", "No Township ID")
		writeResponse(w, http.StatusBadRequest, c.Township)
		return
	}

	err = h.S.SaveCompany(c)
	if err != nil {
		log.Println("[Handler Error SaveCompany]", err)
		writeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeResponse(w, http.StatusCreated, model.ToCompanyReponse(c))
}

func (h *CompanyHandler) UpdateCompany(w http.ResponseWriter, r *http.Request) {
	var c model.CompanyRequest

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		log.Println("[Handler Error UpdateCompany]", err)
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if c.ID == 0 {
		log.Println("[Handler Error UpdateCompany]", "Deep water")
		writeResponse(w, http.StatusInternalServerError, "Information processing error")
		return
	}

	response, err := h.S.UpdateCompany(c)
	if err != nil {
		log.Println("[Handler Error UpdateCompany]", err)
		writeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeResponse(w, http.StatusOK, response)
}

func (h *CompanyHandler) GetLastCompany(w http.ResponseWriter, r *http.Request) {
	var c model.CompanyRequest
	id := h.S.GetLastCompanyID()
	if id == 0 {
		log.Println("[Handler Error GetLastCompany]")
		writeResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	c.ID = id
	writeResponse(w, http.StatusOK, model.ToCompanyReponse(c))
}
