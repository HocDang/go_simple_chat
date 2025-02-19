package ws

import (
	"sync"

	"github.com/google/uuid"
)

type WsHub struct {
	clients    map[uuid.UUID]*WsClient
	broadcast  chan *WsMessage
	register   chan *WsClient
	unregister chan *WsClient
	mu         sync.Mutex
}

type WsMessage struct {
	SenderID   uuid.UUID `json:"sender_id"`
	ReceiverID uuid.UUID `json:"receiver_id"`
	Content    string    `json:"content"`
}

func NewHub() *WsHub {
	return &WsHub{
		clients:    make(map[uuid.UUID]*WsClient),
		broadcast:  make(chan *WsMessage),
		register:   make(chan *WsClient),
		unregister: make(chan *WsClient),
	}
}

func (h *WsHub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client.userID] = client
			h.mu.Unlock()

		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client.userID]; ok {
				delete(h.clients, client.userID)
				close(client.send)
			}
			h.mu.Unlock()

		case message := <-h.broadcast:
			h.mu.Lock()
			receiver, exists := h.clients[message.ReceiverID]
			h.mu.Unlock()
			if exists {
				receiver.send <- []byte(`{"sender_id": "` + message.SenderID.String() + `", "content": "` + message.Content + `"}`)
			}
		}
	}
}
