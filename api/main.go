package main

import (
	"github.com/gmelchert/jogo-da-velha/api/config"
	"github.com/gmelchert/jogo-da-velha/api/routes"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	routes.Initialize()
}
