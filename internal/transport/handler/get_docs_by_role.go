package handler

import (
	"fmt"
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/koyo-os/crm/internal/data/models"
)

func (h *Handler) getByRole(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value("claims").(*models.Claims)
    if !ok {
        http.Error(w, "claims not found", http.StatusInternalServerError)
        return
    }

	docs, err := h.service.GetByRole(claims.UserID, claims.Key)
	if err != nil{
		http.Error(w, "cant get docs", http.StatusInternalServerError)
		return
	}

	resp, err := sonic.Marshal(&docs)
	if err != nil{
		http.Error(w, "cant do resp", http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, resp)
}