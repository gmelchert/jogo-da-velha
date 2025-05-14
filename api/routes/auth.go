package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gmelchert/jogo-da-velha/api/handlers"
	"github.com/gmelchert/jogo-da-velha/api/middleware"
)

func initializeAuthRoutes(router *gin.Engine) {
	public := router.Group("/")

	public.POST("/login", handlers.Login)
	public.POST("/register", handlers.SingUp)

	auth := router.Group("/api/v1")

	auth.Use(middleware.JWTMiddleware())
	auth.GET("/me", handlers.Me)
}
