package handler

import (
	"net/http"

	"github.com/koyo-os/crm/internal/config"
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

func (h *Handler) RegisterRouters(mux *http.ServeMux){
	
}