package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gmelchert/jogo-da-velha/api/models"
	"github.com/gmelchert/jogo-da-velha/api/repository"
	"github.com/gmelchert/jogo-da-velha/api/validator"
)

func FindRoom(ctx *gin.Context) {
	var (
		query             validator.FindRoomQuery
		paginatedResponse models.PaginatedResponse
	)

	if err := ctx.ShouldBindQuery(&query); err != nil {
		Logger.Errorf("error binding query: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, "error binding queries")
		return
	}

	if err := repository.FindRoom(&paginatedResponse, &query); err != nil {
		Logger.Errorf("error finding rooms: %v", err.Error())
		SendError(ctx, http.StatusInternalServerError, "Erro ao encontrar salas")
		return
	}

	SendSuccess(ctx, "find-rooms", http.StatusOK, paginatedResponse)
}

func CreateRoom(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		Logger.Errorf("user unauthneticated error")
		SendError(ctx, http.StatusUnauthorized, "Usuário não autenticado")
		return
	}

	request := validator.CreateRoomRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		Logger.Errorf("error binding request: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, "error binding request")
		return
	}

	if err := request.ValidateCreate(); err != nil {
		Logger.Errorf("validation error: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	payload := validator.CreateRoomPayload{
		OwnerID:    userID.(uint),
		OpponentID: 0,
		Status:     "OPEN",
	}
	room, createErr := repository.CreateRoom(&payload)

	if createErr != nil {
		Logger.Errorf("error creating room: %v", createErr.Error())
		SendError(ctx, http.StatusInternalServerError, "Erro ao criar sala")
		return
	}

	SendSuccess(ctx, "create-room", http.StatusCreated, room)
}

func JoinRoom(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		Logger.Errorf("user unauthneticated error")
		SendError(ctx, http.StatusUnauthorized, "Usuário não autenticado")
		return
	}

	roomID := ctx.Param("id")

	if roomID == "" {
		Logger.Errorf("room id is required")
		SendError(ctx, http.StatusBadRequest, "room id is required")
		return
	}

	if err := repository.JoinRoom(roomID, userID.(uint)); err != nil {
		Logger.Errorf("error joining room: %v", err.Error())
		SendError(ctx, http.StatusInternalServerError, "Erro ao entrar na sala")
		return
	}

	// Disparar evento WS para notificar o jogador que entrou na sala
	// Disparar evento WS para notificar o jogador que criou a sala

	SendSuccess(ctx, "join-room", http.StatusOK, map[string]string{"message": "Você entrou na sala"})
}
