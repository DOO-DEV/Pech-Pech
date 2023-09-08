package main

import (
	"github.com/doo-dev/pech-pech/infrastructure/app"
	"os"
)

func main() {
	app.Run()
	defer os.Exit(0)
}
