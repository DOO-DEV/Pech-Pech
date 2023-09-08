package main

import (
	"github.com/doo-dev/pech-pech/internal/server/adaptor"
	"os"
)

func main() {
	adaptor.NewAdapter()
	defer os.Exit(0)
}
