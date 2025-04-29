package ws

import "sync"

type Hub struct {
	Channels map[string]*Channel
	Mutex    sync.Mutex
}

var GlobalHub = Hub{
	Channels: make(map[string]*Channel),
}

func (h *Hub) GetOrCreateChannel(id string) *Channel {
	h.Mutex.Lock()
	defer h.Mutex.Unlock()

	if channel, ok := h.Channels[id]; ok {
		return channel
	}

	channel := NewChannel(id)
	h.Channels[id] = channel
	return channel
}
