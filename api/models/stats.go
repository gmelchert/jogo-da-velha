package models

import "gorm.io/gorm"

type Stats struct {
	gorm.Model
	UserID      int    `json:"userID"`
	GamesPlayed uint16 `json:"gamesPlayed"`
	Wins        uint16 `json:"wins"`
	Draws       uint16 `json:"draws"`
	Losses      uint16 `json:"losses"`
	User        User   `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
