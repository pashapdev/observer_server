package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/pashapdev/observer_server/internal/application"
	"github.com/pashapdev/observer_server/internal/config"
)

func main() {
	serverCtx, serverStopCtx := context.WithCancel(context.Background())
	app := application.New(config.New())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		app.GracefulStop(serverCtx, sig, serverStopCtx)
	}()

	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}

	<-serverCtx.Done()
}
