package main

import (
	"goevents/config"
	"goevents/internal/app"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	app.Run(cfg)
}
