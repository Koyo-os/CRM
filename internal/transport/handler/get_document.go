package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bytedance/sonic"
	"github.com/koyo-os/crm/internal/data/models"
)

func (h *Handler) getDocument(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value("claims").(*models.Claims)
    if !ok {
        http.Error(w, "claims not found", http.StatusBadRequest)
        return
    }

	id,err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil{
		http.Error(w, "cant get id", http.StatusBadRequest)
		return
	}
	doc, err := h.service.GetDocument(claims.UserID, uint64(id), claims.Key)
	if err != nil{
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	resp, err := sonic.Marshal(doc)
	if err != nil{
		http.Error(w, "cant marshal response", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, resp)
}