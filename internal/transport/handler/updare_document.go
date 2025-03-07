package handler

import (
	"io"
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/koyo-os/crm/internal/data/models"
)

func (h *Handler) UpdateDocument(w http.ResponseWriter, r *http.Request) {
	var docReq models.UpdateReq

	body, err := io.ReadAll(r.Body)
	if err != nil{
		http.Error(w, "cant get req body", http.StatusBadRequest)
		return
	}

	if err = sonic.Unmarshal(body, &docReq);err != nil{
		http.Error(w, "cant unmrashal req", http.StatusBadRequest)
		return
	}

	claims, ok := r.Context().Value("claims").(*models.Claims)
    if !ok {
        http.Error(w, "claims not found", http.StatusBadRequest)
        return
    }

	err = h.service.UpdateDoc(claims.UserID, docReq.ID, claims.Key, &docReq.NewDoc)
	if err != nil{
		http.Error(w, "cant update doc", http.StatusInternalServerError)
		return
	}
}