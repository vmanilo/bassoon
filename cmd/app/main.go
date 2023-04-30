package main

import (
	"context"
	"log"

	"bassoon/config"
	"bassoon/internal/app"
)

var version = "develop"

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalln("failed to load config:", err)
	}

	if err := app.New(version, cfg).Run(context.Background()); err != nil {
		log.Fatal(err)
	}
}
