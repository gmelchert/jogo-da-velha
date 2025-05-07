package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gmelchert/jogo-da-velha/api/models"
	"github.com/gmelchert/jogo-da-velha/api/repositories"
	"github.com/gmelchert/jogo-da-velha/api/utils"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var u models.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := repositories.CreateUser(u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, _ := repositories.AuthenticateUser(u.Username, u.Password)
	token, err := utils.GenerateJWT(user.ID)

	if err != nil {
		http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{"token": token, "username": user.Username, "id": user.ID})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var u models.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := repositories.AuthenticateUser(u.Username, u.Password)
	if err != nil {
		http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{"token": token, "username": user.Username, "id": user.ID})
}

func Me(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int)
	user, _ := repositories.GetUserByID(userID)
	json.NewEncoder(w).Encode(user)
}
