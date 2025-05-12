package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gmelchert/jogo-da-velha/api/handlers"
	"github.com/gmelchert/jogo-da-velha/api/middleware"
)

func initializeRoomRoutes(router *gin.Engine) {
	r := router.Group("/api/v1/rooms")
	r.Use(middleware.JWTMiddleware())

	r.GET("/", handlers.FindRoom)
	r.POST("/", handlers.CreateRoom)
	r.POST("join/:id", handlers.JoinRoom)
}
