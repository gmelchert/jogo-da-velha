package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string `json:"username" gorm:"uniqueIndex"`
	Email       string `json:"email" gorm:"uniqueIndex"`
	Password    string `json:"password"`
	GamesPlayed uint   `json:"gamesPlayed" gorm:"default:0"`
	Wins        uint   `json:"wins" gorm:"default:0"`
	Losses      uint   `json:"losses" gorm:"default:0"`
	Draws       uint   `json:"draws" gorm:"default:0"`
}

type UserResponse struct {
	ID        uint           `json:"id"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
