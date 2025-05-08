package models

import (
	"time"

	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	RoomID     string    `json:"roomId"`
	OwnerID    uint      `json:"userID"`
	Owner      User      `json:"owner" gorm:"foreingKey:UserID"`
	OpponentID uint      `json:"opponentID"`
	Opponent   User      `json:"opponent" gorm:"foreingKey:OpponentID"`
	CreatedAt  time.Time `json:"createdAt"`
}
