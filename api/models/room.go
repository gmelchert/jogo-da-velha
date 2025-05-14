package models

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	RoomID     string `json:"roomId"`
	OwnerID    uint   `json:"ownerId"`
	Owner      User   `json:"owner" gorm:"foreignKey:OwnerID"`
	OpponentID uint   `json:"opponentId"`
	Opponent   User   `json:"opponent" gorm:"foreignKey:OpponentID"`
	Status     string `json:"status"`
}
