package main

import (
	"log"
	"net/http"

	"github.com/gmelchert/jogo-da-velha/api/database"
	"github.com/gmelchert/jogo-da-velha/api/middleware"
	"github.com/gmelchert/jogo-da-velha/api/routes"
)

func main() {
	database.InitDB()

	r := routes.SetupRoutes()

	handler := middleware.CORSMiddleware(r)

	log.Println("Servidor rodando => http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
