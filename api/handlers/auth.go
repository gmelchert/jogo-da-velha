package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gmelchert/jogo-da-velha/api/repository"
	"github.com/gmelchert/jogo-da-velha/api/utils"
	"github.com/gmelchert/jogo-da-velha/api/validator"
	"gorm.io/gorm"
)

type LoginResponse struct {
	Message  string `json:"message"`
	Token    string `json:"token"`
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type MeResponse struct {
	ID        uint           `json:"id"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
	Message   string         `json:"message"`
}

func SingUp(ctx *gin.Context) {
	request := validator.SignUpRequest{}
	ctx.BindJSON(&request)

	if err := request.ValidateSignUp(); err != nil {
		Logger.Errorf("validation error: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, hashPasswordErr := utils.HashPassword(request.Password)
	if hashPasswordErr != nil {
		Logger.Errorf("error hashing password: %v", hashPasswordErr.Error())
		SendError(ctx, http.StatusInternalServerError, "Erro ao registrar usuário")
		return
	}

	request.Password = hashedPassword

	user, CreateErr := repository.CreateUser(&request)
	if CreateErr != nil {
		Logger.Errorf("error creating user: %v", CreateErr.Error())
		SendError(ctx, http.StatusInternalServerError, "Erro ao registrar usuário")
		return
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		SendError(ctx, http.StatusInternalServerError, "Erro ao gerar token")
		return
	}

	response := LoginResponse{
		Message:  "Usuário cadastrado com sucesso",
		Token:    token,
		ID:       user.ID,
		Username: user.Username,
	}

	SendSuccess(ctx, "SignUp", http.StatusCreated, response)
}

func Login(ctx *gin.Context) {
	unauthorizedMessage := "Usuário ou senha incorretos"

	request := validator.LoginRequest{}
	ctx.BindJSON(&request)

	if err := request.ValidateLogin(); err != nil {
		Logger.Errorf("validation error: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user, findUserErr := repository.FindUserByUsername(request.Username)
	if findUserErr != nil {
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
		Message:  "Login realizado com sucesso",
		Token:    token,
		ID:       user.ID,
		Username: user.Username,
	}

	SendSuccess(ctx, "Login", http.StatusOK, response)
}

func Me(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		Logger.Errorf("user unauthenticated error")
		SendError(ctx, http.StatusUnauthorized, "Usuário não autenticado")
		return
	}

	user, err := repository.FindUserByID(userID.(uint))
	if err != nil {
		Logger.Errorf("find user error: %v", err.Error())
		SendError(ctx, http.StatusInternalServerError, "Usuário não encontrado")
		return
	}

	response := MeResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
		Message:   "Usuário encontrado com sucesso",
	}

	SendSuccess(ctx, "Me", http.StatusOK, response)
}
