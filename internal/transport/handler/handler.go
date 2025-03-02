package handler

import (
	"github.com/koyo-os/crm/internal/config"
	"github.com/koyo-os/crm/internal/service"
	"github.com/koyo-os/crm/pkg/loger"
)

type Handler struct{
	service *service.Service
	loger loger.Logger
}

func New(cfg *config.Config) (*Handler, error) {
	
}