package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gmelchert/jogo-da-velha/api/handlers"
	"github.com/gmelchert/jogo-da-velha/api/utils"
)

func JWTMiddleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			http.Error(w, "Token não fornecido", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(auth, "Bearer ")
		userID, err := utils.ParseToken(token)
		if err != nil {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")

		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			handlers.SendError(c, http.StatusUnauthorized, "Token não fornecido")
		}

		token := strings.TrimPrefix(auth, "Bearer ")
		userID, err := utils.ParseToken(token)
		if err != nil {
			handlers.SendError(c, http.StatusUnauthorized, "Token inválido")
			return
		}

		c.Set("userID", userID)

		c.Next()
	}
}
