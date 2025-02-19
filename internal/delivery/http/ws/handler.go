package ws

import (
	"chat-server/internal/domain/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type websocketHandler struct {
	hub            *WsHub
	messageUseCase usecases.MessageUseCase
}

func NewWebSocketHandler(hub *WsHub, messageUseCase usecases.MessageUseCase) *websocketHandler {
	return &websocketHandler{hub: hub, messageUseCase: messageUseCase}
}

var wsUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func (h *websocketHandler) WebSocketHandler(c *gin.Context) {
	userID := c.MustGet("x-user-id").(uuid.UUID)

	conn, err := wsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	client := NewClient(h.hub, conn, userID, h.messageUseCase)
	h.hub.register <- client

	go client.ReadMessages()
	go client.WriteMessages()
}
