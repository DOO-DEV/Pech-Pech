package api

import (
	"context"
	"fmt"
	"github.com/doo-dev/pech-pech/infrastructure/mail"
	autConfing "github.com/doo-dev/pech-pech/internal/modules/auth/usecase"
	"github.com/doo-dev/pech-pech/pkg/constants"
	"github.com/doo-dev/pech-pech/pkg/richerror"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type Config struct {
	Port                     string        `koanf:"port"`
	ReadTimeoutInSeconds     time.Duration `koanf:"read_timeout_in_seconds"`
	WriteTimeoutInSeconds    time.Duration `koanf:"write_timeout_in_seconds"`
	ShutDownTimeoutInSeconds time.Duration `koanf:"shut_down_timeout_in_seconds"`
}

type Api struct {
	Echo     *echo.Echo
	pgDB     *gorm.DB
	authConf autConfing.Config
	mailConf mail.Config
	cfg      Config
}

type IApi interface {
	HttpApi() error
	Start(chan error)
	Stop(context.Context) error
}

func NewApi(cfg Config, authCfg autConfing.Config, mailCfg mail.Config, e *echo.Echo, pgDB *gorm.DB) IApi {
	return Api{
		Echo:     e,
		pgDB:     pgDB,
		authConf: authCfg,
		mailConf: mailCfg,
		cfg:      cfg,
	}
}

func (a Api) Start(c chan error) {
	const op = "api.Start"

	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%s", a.cfg.Port),
		WriteTimeout: time.Second * a.cfg.WriteTimeoutInSeconds,
		ReadTimeout:  time.Second * a.cfg.ReadTimeoutInSeconds,
	}
	if err := a.HttpApi(); err != nil {
		c <- richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(constants.ErrMsgSetupHttpRouter)
	}

	go func() {
		if err := a.Echo.StartServer(httpServer); err != nil {
			c <- richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected).WithMessage(constants.ErrMsgStartHttp)
		}
	}()
	
}

func (a Api) Stop(ctx context.Context) error {
	return a.Echo.Shutdown(ctx)
}
