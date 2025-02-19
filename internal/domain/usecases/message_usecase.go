package usecases

import (
	"chat-server/internal/domain/entities"
)

type MessageUseCase interface {
	SendMessage(senderID, receiverID int, content string) error
	GetMessages(receiverID int) ([]entities.Message, error)
}
