package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/koyo-os/crm/internal/config"
	"github.com/koyo-os/crm/internal/data/models"
	"github.com/koyo-os/crm/internal/service"
	"github.com/koyo-os/crm/pkg/loger"
)

type Handler struct{
	service *service.Service
	loger loger.Logger
}

func New(cfg *config.Config) (*Handler, error) {
	service, err := service.New(cfg)
	return &Handler{
		service: service,
		loger: loger.New(),
	}, err
}

func (h *Handler) AddDocument(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value("claims").(*models.Claims)
    if !ok {
        http.Error(w, "claims not found", http.StatusInternalServerError)
        return
    }

	var document models.Document

	body, err := io.ReadAll(r.Body)
	if err != nil{
		http.Error(w, "cant get req body", http.StatusBadRequest)
		return
	}

	if err = sonic.Unmarshal(body, &document);err != nil{
		http.Error(w, "cant unmarshal document", http.StatusBadRequest)
		return
	}

	id, err := h.service.AddDocument(uint64(claims.UserID), claims.Key, &document)
	if err != nil{
		http.Error(w, "cant add document", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, map[string]string{"id" : fmt.Sprintf("%d", id)})
}