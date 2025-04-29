package ws

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID       int
	Conn     *websocket.Conn
	Send     chan []byte
	Channel  *Channel
	Username string
}

func (c *Client) ReadLoop() {
	defer func() {
		c.Channel.RemovePlayer(c)
		c.Conn.Close()
	}()

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("Erro ao ler:", err)
			break
		}

		c.Channel.Broadcast <- msg
	}
}

func (c *Client) WriteLoop() {
	defer c.Conn.Close()

	for msg := range c.Send {
		if err := c.Conn.WriteMessage(1, msg); err != nil {
			log.Println("Erro ao escrever:", err)
			break
		}
	}
}
