package repositories

import (
	"chat-server/internal/domain/entities"

	"github.com/google/uuid"
)

type MessageRepository interface {
	Create(message *entities.Message) error
	GetByReceiverID(receiverID uuid.UUID, senderID uuid.UUID) ([]entities.Message, error)
}
