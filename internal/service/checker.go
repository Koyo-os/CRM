package service

import (
	"github.com/koyo-os/crm/internal/config"
	"github.com/koyo-os/crm/pkg/loger"
)

type Checker struct{
	service *Service
	logger loger.Logger
}

func NewChecker(cfg *config.Config) (*Checker, error){
	service, err := New(cfg)
	if err != nil{
		return nil,err
	}
	return &Checker{
		service: service,
		logger: loger.New(),
	}, nil
}

func (c *Checker) Check() {
	c.logger.Info().Msg("starting role check")

	var ch chan error
	go c.service.CheckAllUserRoleTimes(ch)

	err := <- ch
	if err != nil{
		c.logger.Error().Err(err)
	}

	c.logger.Info().Msg("success check!")
}