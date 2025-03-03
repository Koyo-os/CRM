package handler

import (
	"net/http"
	"strconv"

	"github.com/koyo-os/crm/internal/data/models"
)

func (h *Handler) deleteDocument(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value("claims").(*models.Claims)
    if !ok {
        http.Error(w, "claims not found", http.StatusInternalServerError)
        return
    }
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil{
		http.Error(w, "cant get id", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteDocument(uint64(id), claims.UserID, claims.Key)
	if err != nil{
		http.Error(w, "cant delete document", http.StatusInternalServerError)
		return
	}
}