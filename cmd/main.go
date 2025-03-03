package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/koyo-os/crm/internal/app"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	app.Init().Run(ctx)
}