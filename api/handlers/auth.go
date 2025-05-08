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

func SingUp(ctx *gin.Context) {
	request := models.SignUpRequest{}
	ctx.BindJSON(&request)

	if err := request.ValidateSignUp(); err != nil {
		Logger.Errorf("validation error: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
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

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		SendError(ctx, http.StatusInternalServerError, "Erro ao gerar token")
		return
	}

	response := LoginResponse{
		Message: "Usuário cadastrado com sucesso",
		Token:   token,
		ID:      user.ID,
	}

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

	err := Db.First(&user, "username = ?", request.Username)
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

func Me1(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int)
	user, _ := repositories.GetUserByID(userID)
	json.NewEncoder(w).Encode(user)
}

func Me(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		SendError(ctx, http.StatusUnauthorized, "Usuário não autenticado")
		return
	}

	user := models.User{}

	if err := Db.First(&user, userID); err != nil {
		SendError(ctx, http.StatusInternalServerError, "Usuário não autenticado")
		return
	}

}
