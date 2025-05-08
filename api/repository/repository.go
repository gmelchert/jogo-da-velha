package repository

import (
	"github.com/gmelchert/jogo-da-velha/api/config"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

func InitializeRepository() {
	Db = config.GetSQLite()
}
