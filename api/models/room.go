package models

import (
	"time"

	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	ID         string    `json:"id"`
	OwnerID    int       `json:"userID"`
	Owner      User      `json:"owner" gorm:"foreingKey:UserID"`
	OpponentID int       `json:"opponentID"`
	Opponent   User      `json:"opponent" gorm:"foreingKey:OpponentID"`
	CreatedAt  time.Time `json:"createdAt"`
}
