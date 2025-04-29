package ws

type Message struct {
	Type string      `json:"json"`
	Data interface{} `json:"data"`
}

type PlayData struct {
	Position int `json:"position"`
}
