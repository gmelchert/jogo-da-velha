package repository

import (
	"fmt"

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

// operation: 1 = win, 2 = loss, 3 = draw
func UpdateUserStat(id uint, op uint8) (models.User, error) {
	var (
		user models.User
		err  error
	)

	if op != 1 && op != 2 && op != 3 {
		return user, fmt.Errorf("invalid operation: %d", op)
	}

	user, err = FindUserByID(id)
	if err != nil {
		return user, err
	}

	wins := user.Wins
	losses := user.Losses
	draws := user.Draws

	switch op {
	case 1:
		wins = wins + 1
	case 2:
		losses = losses + 1
	case 3:
		draws = draws + 1
	}

	if err := Db.Model(&user).Where("id = ?", id).Updates(models.User{
		GamesPlayed: user.GamesPlayed + 1,
		Wins:        wins,
		Losses:      losses,
		Draws:       draws,
	}).Error; err != nil {
		return user, err
	}

	return user, nil
}
