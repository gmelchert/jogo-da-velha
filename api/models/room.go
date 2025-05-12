package models

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	RoomID     string `json:"roomId"`
	OwnerID    uint   `json:"userId"`
	Owner      User   `json:"owner" gorm:"foreingKey:UserID"`
	OpponentID uint   `json:"opponentID"`
	Opponent   User   `json:"opponent" gorm:"foreingKey:OpponentID"`
	Status     string `json:"status"`
}
