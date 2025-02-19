package http

import (
	"chat-server/internal/domain/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MessageHandler struct {
	messageUseCase usecases.MessageUseCase
}

func NewMessageHandler(messageUseCase usecases.MessageUseCase) *MessageHandler {
	return &MessageHandler{messageUseCase: messageUseCase}
}

// Gửi tin nhắn
func (h *MessageHandler) SendMessage(c *gin.Context) {
	var req struct {
		SenderID   uuid.UUID `json:"sender_id" binding:"required"`
		ReceiverID uuid.UUID `json:"receiver_id" binding:"required"`
		Content    string    `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.messageUseCase.SendMessage(req.SenderID, req.ReceiverID, req.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}

// Lấy tin nhắn của người nhận
func (h *MessageHandler) GetMessages(c *gin.Context) {
	receiverIDStr := c.Param("receiver_id")
	receiverID, err := uuid.Parse(receiverIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receiver ID"})
		return
	}

	messages, err := h.messageUseCase.GetMessages(receiverID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get messages"})
		return
	}

	c.JSON(http.StatusOK, messages)
}
