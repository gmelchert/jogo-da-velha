package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gmelchert/jogo-da-velha/api/handlers"
	"github.com/gmelchert/jogo-da-velha/api/middleware"
	"github.com/gmelchert/jogo-da-velha/api/repository"
)

func Initialize() {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	handlers.InitializeHandler()
	repository.InitializeRepository()

	initializeAuthRoutes(r)
	initializeRoomRoutes(r)

	r.Run(":8080")
}
