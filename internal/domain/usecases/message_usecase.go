package usecases

import (
	"chat-server/internal/domain/entities"

	"github.com/google/uuid"
)

type MessageUseCase interface {
	SendMessage(senderID uuid.UUID, receiverID uuid.UUID, content string) error
	GetMessages(receiverID uuid.UUID) ([]entities.Message, error)
}
