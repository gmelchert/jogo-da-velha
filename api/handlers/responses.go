package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func SendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")

	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func SendSuccess(ctx *gin.Context, op string, status int, data interface{}) {
	ctx.Header("Content-type", "application/json")

	ctx.JSON(status, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successful", op),
		"data":    data,
	})
}
