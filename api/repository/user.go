package repository

import (
	"github.com/gmelchert/jogo-da-velha/api/models"
	"github.com/gmelchert/jogo-da-velha/api/validator"
)

func FindUserByEmail(email string) (models.User, error) {
	var user models.User
	if err := Db.Where("email = ?", email).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func FindUserByID(id uint) (models.User, error) {
	user := models.User{}

	if err := Db.First(&user, id).Error; err != nil {
		return user, err
	}

	return user, nil
}

func FindUserByUsername(username string) (models.User, error) {
	var user models.User
	if err := Db.Where("username = ?", username).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func CreateUser(req *validator.SignUpRequest) (models.User, error) {
	user := models.User{
		Username:    req.Username,
		Email:       req.Email,
		Password:    req.Password,
		GamesPlayed: 0,
		Wins:        0,
		Losses:      0,
		Draws:       0,
	}

	if err := Db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
