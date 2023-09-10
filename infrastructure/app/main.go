package app

import (
	"context"
	"github.com/doo-dev/pech-pech/config"
	"github.com/doo-dev/pech-pech/internal/server/adaptor"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	cfg := config.Load()

	c := make(chan os.Signal, 1)

	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGHUP, syscall.SIGQUIT)
	signal.Notify(c, os.Kill)

	adt := adaptor.NewAdapter(cfg.PgDB, cfg.HttpServer, cfg.AuthConfig, cfg.MailConfig)

	select {
	case s := <-c:
		log.Printf("\nserver got terminate: %s\n", s.String())
	case err := <-adt.Notify():
		log.Printf("\nserver go error: %s\n", err.Error())
	}

	// TODO - config timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*cfg.HttpServer.ShutDownTimeoutInSeconds)
	defer cancel()

	p, _ := os.FindProcess(os.Getpid())
	p.Signal(syscall.SIGINT)

	if err := adt.Shutdown(ctx); err != nil {
		log.Printf("%v\n", err)
	}
}
