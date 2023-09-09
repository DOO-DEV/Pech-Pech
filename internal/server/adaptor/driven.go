package adaptor

import (
	"context"
	"github.com/doo-dev/pech-pech/db/postgres"
	"github.com/doo-dev/pech-pech/infrastructure/mail"
	"github.com/doo-dev/pech-pech/internal/modules/auth/usecase"
	"github.com/doo-dev/pech-pech/internal/server/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type IAdapter interface {
	Notify() <-chan error
	Shutdown(ctx context.Context) error
}

type Adapter struct {
	Api api.IApi
}

var signal = make(chan error, 1)

func NewAdapter(pgCfg postgres.Config, httpCfg api.Config, authCfg usecase.Config, mailCfg mail.Config) IAdapter {
	pgAdaptor := postgres.NewPostgresAdapter(pgCfg)
	pgInstance := pgAdaptor.ConnectInstance()

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	initServer := api.NewApi(httpCfg, authCfg, mailCfg, e, pgInstance)
	initServer.Start(signal)

	return Adapter{Api: initServer}
}

func (a Adapter) Notify() <-chan error {
	return signal
}

func (a Adapter) Shutdown(ctx context.Context) error {
	return a.Api.Stop(ctx)
}
