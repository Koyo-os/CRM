package handler

import (
	"net/http"

	"github.com/koyo-os/crm/internal/data/models"
)

func (h *Handler) getByRole(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value("claims").(*models.Claims)
    if !ok {
        http.Error(w, "claims not found", http.StatusInternalServerError)
        return
    }

	
}