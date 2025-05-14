package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gmelchert/jogo-da-velha/api/handlers"
	"github.com/gmelchert/jogo-da-velha/api/utils"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		auth := c.Request.Header.Get("Authorization")

		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			handlers.SendError(c, http.StatusUnauthorized, "Token não fornecido")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(auth, "Bearer ")
		userID, err := utils.ParseToken(token)
		if err != nil {
			handlers.SendError(c, http.StatusUnauthorized, "Token inválido")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("userID", uint(userID))

		c.Next()
	}
}
