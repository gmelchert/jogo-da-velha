package models

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	RoomID     string `json:"roomId"`
	OwnerID    uint   `json:"ownerID"`
	Owner      User   `json:"owner" gorm:"foreingKey:OwnerID"`
	OpponentID uint   `json:"opponentID"`
	Opponent   User   `json:"opponent" gorm:"foreingKey:OpponentID"`
	Status     string `json:"status"`
}
