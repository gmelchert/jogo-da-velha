package handlers

import (
	"fmt"
	"net/http"

	"github.com/gmelchert/jogo-da-velha/api/utils"
	"github.com/gmelchert/jogo-da-velha/api/ws"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	channelID := mux.Vars(r)["channelID"]

	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "Token ausente", http.StatusUnauthorized)
		return
	}

	userID, _ := utils.ParseToken(token)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "WebSocket upgrade failed", http.StatusBadRequest)
		return
	}

	client := &ws.Client{
		ID:   userID,
		Conn: conn,
		Send: make(chan []byte),
	}

	channel := ws.GlobalHub.GetOrCreateChannel(channelID)
	client.Channel = channel
	channel.AddPlayer(client)

	go client.ReadLoop()
	go client.WriteLoop()

	go channel.Run()

	fmt.Println("Jogador conectado na sala", channelID)
}
