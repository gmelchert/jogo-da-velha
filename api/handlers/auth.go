package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gmelchert/jogo-da-velha/api/models"
	"github.com/gmelchert/jogo-da-velha/api/repositories"
	"github.com/gmelchert/jogo-da-velha/api/utils"
)

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
	ID      uint   `json:"id"`
}

func SignUp1(w http.ResponseWriter, r *http.Request) {
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

func SingUp(ctx *gin.Context) {
	request := models.SignUpRequest{}
	ctx.BindJSON(&request)

	if err := request.ValidateSignUp(); err != nil {
		Logger.Errorf("validation error: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var count int

	Db.Model(models.User).Count(&count) // corrigir
	if count > 0 {
		SendError(ctx, http.StatusUnauthorized, "Nome ou email do usuário já estão cadastrados")
		return
	}

	user := models.User{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}

	err := Db.Create(&user).Error
	if err != nil {
		Logger.Errorf("error creating user: %v", err.Error())
		SendError(ctx, http.StatusInternalServerError, "Erro ao registrar usuário")
		return
	}

	response := LoginResponse{}

	SendSuccess(ctx, "signUp", http.StatusCreated, response)
}

func Login(ctx *gin.Context) {
	request := models.LoginRequest{}
	ctx.BindJSON(&request)

	if err := request.ValidateLogin(); err != nil {
		Logger.Errorf("validation error: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := models.User{}
	unauthorizedMessage := "Usuário ou senha incorretos"

	err := Db.Limit(1).First(&user, "username = ?", request.Username)
	if err != nil {
		SendError(ctx, http.StatusUnauthorized, unauthorizedMessage)
		return
	}

	if verifyPassword := utils.VerifyPassword(request.Password, user.Password); !verifyPassword {
		SendError(ctx, http.StatusUnauthorized, unauthorizedMessage)
		return
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		SendError(ctx, http.StatusInternalServerError, "Erro ao gerar token")
		return
	}

	response := LoginResponse{
		Message: "Login realizado com sucesso",
		Token:   token,
		ID:      user.ID,
	}

	SendSuccess(ctx, "login", http.StatusOK, response)
}

func Me(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int)
	user, _ := repositories.GetUserByID(userID)
	json.NewEncoder(w).Encode(user)
}
