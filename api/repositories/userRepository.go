package repositories

import (
	"errors"

	"github.com/gmelchert/jogo-da-velha/api/database"
	"github.com/gmelchert/jogo-da-velha/api/models"
)

func CreateUser(u models.User) error {
	_, err := database.DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", u.Username, u.Password)
	return err
}

func AuthenticateUser(username, password string) (models.User, error) {
	row := database.DB.QueryRow("SELECT id, username FROM users WHERE username = ? AND password = ?", username, password)
	var user models.User

	if err := row.Scan(&user.ID, &user.Username); err != nil {
		return user, errors.New("usuário ou senha inválidos")
	}

	return user, nil
}

func GetUserByID(id int) (models.User, error) {
	row := database.DB.QueryRow("SELECT id, username FROM users WHERE id = ?", id)
	var user models.User

	err := row.Scan(&user.ID, &user.Username)
	return user, err
}
