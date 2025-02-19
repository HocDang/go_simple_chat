package ws

import (
	"chat-server/internal/domain/usecases"
	"log"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type WsClient struct {
	hub        *WsHub
	conn       *websocket.Conn
	send       chan []byte
	userID     uuid.UUID
	messageUse usecases.MessageUseCase
}

func NewClient(hub *WsHub, conn *websocket.Conn, userID uuid.UUID, messageUse usecases.MessageUseCase) *WsClient {
	return &WsClient{
		hub:        hub,
		conn:       conn,
		send:       make(chan []byte, 256),
		userID:     userID,
		messageUse: messageUse,
	}
}

func (c *WsClient) ReadMessages() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	for {
		var msg WsMessage
		err := c.conn.ReadJSON(&msg)
		if err != nil {
			log.Println("❌ Read error:", err)
			break
		}
		msg.SenderID = c.userID
		err = c.messageUse.SendMessage(msg.SenderID, msg.ReceiverID, msg.Content)
		if err != nil {
			log.Println("❌ Failed to send message:", err)
			break
		}
		c.hub.broadcast <- &msg
	}
}

func (c *WsClient) WriteMessages() {
	defer c.conn.Close()
	for message := range c.send {
		err := c.conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("❌ Write error:", err)
			break
		}
	}
}
