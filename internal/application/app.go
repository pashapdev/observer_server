package application

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/pashapdev/observer_server/internal/config"
	"github.com/pashapdev/observer_server/internal/router"
)

type App struct {
	conf   *config.Config
	server *http.Server
}

type AppOption func(*App)

func New(conf *config.Config, opts ...AppOption) *App {
	app := new(App)
	app.conf = conf
	app.server = &http.Server{
		Addr:        conf.Address,
		Handler:     router.New(),
		ReadTimeout: 5 * time.Minute,
	}
	return app
}

func (a *App) Run() error {
	log.Println("starting http server...")
	err := a.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Println("server was stop with err", err)
		return err
	}
	log.Println("server was stop")
	return nil
}

func (a *App) stop(ctx context.Context) error {
	log.Println("shutdowning server...")
	err := a.server.Shutdown(ctx)
	if err != nil {
		log.Println("server was shutdown with error", err)
		return err
	}
	log.Println("server was shutdown")
	return nil
}

func (a *App) GracefulStop(serverCtx context.Context, sig <-chan os.Signal, serverStopCtx context.CancelFunc) {
	<-sig
	var timeOut = 30 * time.Second
	shutdownCtx, shutdownStopCtx := context.WithTimeout(serverCtx, timeOut)

	go func() {
		<-shutdownCtx.Done()
		if shutdownCtx.Err() == context.DeadlineExceeded {
			log.Fatal("graceful shutdown timed out... forcing exit.")
		}
	}()

	err := a.stop(shutdownCtx)
	if err != nil {
		log.Fatal(err)
	}
	serverStopCtx()
	shutdownStopCtx()
}
