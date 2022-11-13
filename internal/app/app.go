package app

import (
	"context"

	"github.com/pashapdev/observer_server/internal/config"
)

type App struct {
}

type AppOption func(*App)

func New(ctx context.Context, conf *config.Config, opts ...AppOption) (*App, error) {
	application := new(App)
	return application, nil
}
