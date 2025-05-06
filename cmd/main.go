package main

import (
	"github.com/Tairascii/donation-managment-system/internal"
	"github.com/Tairascii/donation-managment-system/internal/infra"
)

func main() {
	cfg, err := infra.LoadConfig()
	if err != nil {
		panic(err)
	}
	app := internal.New(cfg)
	if err = app.Run(); err != nil {
		panic(err)
	}
}
