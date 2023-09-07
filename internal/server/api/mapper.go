package api

import (
	"context"
	"github.com/doo-dev/pech-pech/pkg/constants"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type Api struct {
	Echo *echo.Echo
	pgDB *gorm.DB
}

type IApi interface {
	HttpApi() error
	Start(chan error)
	Stop(context.Context) error
}

func NewApi(e *echo.Echo, pgDB *gorm.DB) IApi {
	return Api{
		Echo: e,
		pgDB: pgDB,
	}
}

func (a Api) Start(c chan error) {
	httpServer := &http.Server{
		Addr:         ":8080",
		WriteTimeout: time.Second * 60,
		ReadTimeout:  time.Second * 60,
	}
	if err := a.HttpApi(); err != nil {
		c <- constants.ErrSetupHttpRouter
	}

	if err := a.Echo.StartServer(httpServer); err != nil {
		c <- constants.ErrStartHttp
	}
}

func (a Api) Stop(ctx context.Context) error {
	return a.Echo.Shutdown(ctx)
}
