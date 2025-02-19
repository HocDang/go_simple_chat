package usecase

import (
	"chat-server/internal/domain/entities"
	"chat-server/internal/domain/repositories"

	"github.com/google/uuid"
)

type MessageUseCase struct {
	messageRepo repositories.MessageRepository
}

func NewMessageUseCase(messageRepo repositories.MessageRepository) *MessageUseCase {
	return &MessageUseCase{messageRepo: messageRepo}
}

func (uc *MessageUseCase) SendMessage(senderID uuid.UUID, receiverID uuid.UUID, content string) error {
	message := &entities.Message{
		SenderID:   senderID,
		ReceiverID: receiverID,
		Content:    content,
	}

	return uc.messageRepo.Create(message)
}

func (uc *MessageUseCase) GetMessages(receiverID uuid.UUID) ([]entities.Message, error) {
	return uc.messageRepo.GetByReceiverID(receiverID)
}
