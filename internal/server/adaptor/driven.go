package adaptor

import (
	"context"
	"github.com/doo-dev/pech-pech/db/postgres"
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

func NewAdapter() IAdapter {
	pgAdaptor := postgres.NewPostgresAdapter()
	pgInstance := pgAdaptor.ConnectInstance()

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	initServer := api.NewApi(e, pgInstance)
	initServer.Start(signal)

	return Adapter{Api: initServer}
}

func (a Adapter) Notify() <-chan error {
	return signal
}

func (a Adapter) Shutdown(ctx context.Context) error {
	return a.Api.Stop(ctx)
}
