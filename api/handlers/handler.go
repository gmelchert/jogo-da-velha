package handlers

import (
	"github.com/gmelchert/jogo-da-velha/api/config"
)

var (
	Logger *config.Logger
)

func InitializeHandler() {
	Logger = config.GetLogger("handler")
}
