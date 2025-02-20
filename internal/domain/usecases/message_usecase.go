package usecases

import (
	"chat-server/internal/domain/entities"

	"github.com/google/uuid"
)

type MessageUseCase interface {
	SendMessage(senderID uuid.UUID, receiverID uuid.UUID, content string) error
	GetMessage(receiverID uuid.UUID, senderID uuid.UUID) ([]entities.Message, error)
	SearchMessage(receiverID uuid.UUID, senderID uuid.UUID, query string) ([]entities.Message, error)
}
