package app

import (
	"context"
	"github.com/doo-dev/pech-pech/internal/server/adaptor"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	c := make(chan os.Signal, 1)

	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGHUP, syscall.SIGQUIT)
	signal.Notify(c, os.Kill)

	adt := adaptor.NewAdapter()

	select {
	case s := <-c:
		log.Printf("\nserver got terminate: %s\n", s.String())
	case err := <-adt.Notify():
		log.Printf("\nserver go error: %s\n", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	p, _ := os.FindProcess(os.Getpid())
	p.Signal(syscall.SIGINT)

	if err := adt.Shutdown(ctx); err != nil {
		log.Printf("%v\n", err)
	}
}
