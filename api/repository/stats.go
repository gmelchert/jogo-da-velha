package repository

import "github.com/gmelchert/jogo-da-velha/api/models"

func FindStatsByUserID(userID uint) (models.Stats, error) {
	stats := models.Stats{}
	if err := Db.Where("user_id = ?", userID).First(&stats).Error; err != nil {
		return stats, err
	}
	return stats, nil
}
