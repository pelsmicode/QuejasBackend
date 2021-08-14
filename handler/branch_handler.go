package handler

import (
	"log"
	"net/http"

	"github.com/pelsmicode/QuejasBackend/service"
)

type BranchHandler struct {
	S service.BranchService
}

func (h *BranchHandler) GetDiacoBranches(w http.ResponseWriter, r *http.Request) {
	branches, err := h.S.GetBranches()
	if err != nil {
		log.Println("[Handler Error GetBranches]", err)
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	writeResponse(w, http.StatusOK, branches)
}
