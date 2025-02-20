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

func (h *MessageHandler) SendMessage(c *gin.Context) {
	var req struct {
		ReceiverID uuid.UUID `json:"receiver_id" binding:"required"`
		Content    string    `json:"content" binding:"required"`
	}
	senderID := c.MustGet("x-user-id").(uuid.UUID)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.messageUseCase.SendMessage(senderID, req.ReceiverID, req.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}

func (h *MessageHandler) GetMessages(c *gin.Context) {
	receiverIDStr := c.Param("id")
	receiverID, err := uuid.Parse(receiverIDStr)
	senderID := c.MustGet("x-user-id").(uuid.UUID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receiver ID"})
		return
	}

	messages, err := h.messageUseCase.GetMessage(receiverID, senderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get messages"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": messages})
}

func (h *MessageHandler) SearchMessages(c *gin.Context) {
	receiverIDStr := c.Param("id")
	receiverID, err := uuid.Parse(receiverIDStr)
	senderID := c.MustGet("x-user-id").(uuid.UUID)
	keyword := c.Query("keyword")

	if keyword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Keyword is required"})
		return
	}

	results, err := h.messageUseCase.SearchMessage(receiverID, senderID, keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search messages"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": results})
}
