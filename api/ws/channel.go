package ws

import (
	"encoding/json"
	"log"
	"sync"
)

type Channel struct {
	ID        string
	Players   []*Client
	Broadcast chan []byte
	Mutex     sync.Mutex

	Board      [9]int
	Turn       int
	ActiveGame bool
}

var winningConditions = [][3]int{
	{0, 1, 2}, {3, 4, 5}, {6, 7, 8},
	{0, 3, 6}, {1, 4, 7}, {2, 5, 8},
	{0, 4, 8}, {2, 4, 6},
}

func checkWinCondition(tab [9]int, player int) bool {
	for _, c := range winningConditions {
		if tab[c[0]] == player && tab[c[1]] == player && tab[c[2]] == player {
			return true
		}
	}
	return false
}

func NewChannel(id string) *Channel {
	return &Channel{
		ID:        id,
		Players:   make([]*Client, 0, 2),
		Broadcast: make(chan []byte),
	}
}

func (ch *Channel) AddPlayer(cl *Client) {
	ch.Mutex.Lock()
	defer ch.Mutex.Unlock()

	if len(ch.Players) < 2 {
		ch.Players = append(ch.Players, cl)
	}

	if len(ch.Players) == 2 && !ch.ActiveGame {
		ch.Turn = 1
		ch.Board = [9]int{}
		ch.ActiveGame = true

		ch.sendToAll("start", map[string]int{"turn": ch.Turn})
	}
}

func (ch *Channel) RemovePlayer(cl *Client) {
	ch.Mutex.Lock()
	defer ch.Mutex.Unlock()

	for i, player := range ch.Players {
		if player == cl {
			ch.Players = append(ch.Players[:i], ch.Players[:i+1]...)
			break
		}
	}
}

func (c *Channel) Run() {
	for {
		msg := <-c.Broadcast

		var input Message
		if err := json.Unmarshal(msg, &input); err != nil {
			log.Println("Mensagem inválida:", err)
			continue
		}

		switch input.Type {
		case "play":
			dataMap, _ := input.Data.(map[string]interface{})
			pos := int(dataMap["position"].(float64))
			c.processPlay(pos)
		}
	}
}

func (c *Channel) processPlay(pos int) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	if !c.ActiveGame || pos < 0 || pos > 8 || c.Board[pos] != 0 {
		c.sendToAll("invalid", "Jogada inválida")
		return
	}

	c.Board[pos] = c.Turn

	c.sendToAll("play", map[string]interface{}{
		"position": pos,
		"player":   c.Turn,
	})

	if checkWinCondition(c.Board, c.Turn) {
		c.sendToAll("draw", nil)
		c.ActiveGame = false
		return
	}

	if c.Turn == 1 {
		c.Turn = 2
	} else {
		c.Turn = 1
	}
}

func isDraw(tab [9]int) bool {
	for _, v := range tab {
		if v == 0 {
			return false
		}
	}
	return true
}

func (c *Channel) sendToAll(msgType string, data interface{}) {
	msg := Message{
		Type: msgType,
		Data: data,
	}
	jsonMsg, _ := json.Marshal(msg)

	for _, player := range c.Players {
		player.Send <- jsonMsg
	}
}
