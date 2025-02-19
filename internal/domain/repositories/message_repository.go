package repositories

import "chat-server/internal/domain/entities"

type MessageRepository interface {
	Create(message *entities.Message) error
	GetByReceiverID(receiverID int) ([]entities.Message, error)
}
