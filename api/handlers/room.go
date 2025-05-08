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

	room, createErr := repository.CreateRoom(&request)

	if createErr != nil {
		Logger.Errorf("error creating room: %v", createErr.Error())
		SendError(ctx, http.StatusInternalServerError, "Erro ao criar sala")
		return
	}

	SendSuccess(ctx, "create-room", http.StatusCreated, room)
}
