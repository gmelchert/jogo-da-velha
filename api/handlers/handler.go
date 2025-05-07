package handlers

import (
	"github.com/gmelchert/jogo-da-velha/api/config"
	"gorm.io/gorm"
)

var (
	Logger *config.Logger
	Db     *gorm.DB
)

func InitializeHandler() {
	Logger = config.GetLogger("handler")
	Db = config.GetSQLite()
}
