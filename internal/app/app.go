package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/koyo-os/crm/internal/config"
	"github.com/koyo-os/crm/internal/transport/handler"
	"github.com/koyo-os/crm/pkg/loger"
)

type App struct{
	loger loger.Logger
	handler *handler.Handler
	cfg *config.Config
	server *http.Server
}

func Init() *App {
	logger := loger.New()
	cfg := config.Load()
	handler, err := handler.New(cfg)
	if err != nil{
		logger.Error().Err(err)
		return nil
	}

	return &App{
		loger: logger,
		cfg: cfg,
		handler: handler,
		server: &http.Server{
			Addr: fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		},
	}
}

func (a *App) Run(ctx context.Context) {
	go func ()  {
		<- ctx.Done()
		a.server.Shutdown(ctx)
		a.loger.Info().Msg("server stopped")	
	}()

	mux := http.NewServeMux()

	a.loger.Info().Msg("register routing...")
	a.handler.RegisterRouters(mux)
}