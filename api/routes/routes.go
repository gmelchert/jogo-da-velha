package routes

import (
	"github.com/gmelchert/jogo-da-velha/api/handlers"
	"github.com/gmelchert/jogo-da-velha/api/middleware"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/signup", handlers.SignUp).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	// rota protegida
	s := r.PathPrefix("/api").Subrouter()
	s.Use(middleware.JWTMiddleware)

	s.HandleFunc("/auth/me", handlers.Me).Methods("GET")

	s.HandleFunc("/ws/{channelID}", handlers.WebSocketHandler) // .Methods("GET")

	return r
}
