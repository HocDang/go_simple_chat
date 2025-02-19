package usecase

import (
	"chat-server/internal/domain/entities"
	"chat-server/internal/domain/repositories"
	"time"
)

type MessageUseCase struct {
	messageRepo repositories.MessageRepository
}

func NewMessageUseCase(messageRepo repositories.MessageRepository) *MessageUseCase {
	return &MessageUseCase{messageRepo: messageRepo}
}

func (uc *MessageUseCase) SendMessage(senderID, receiverID int, content string) error {
	message := &entities.Message{
		SenderID:   senderID,
		ReceiverID: receiverID,
		Content:    content,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	return uc.messageRepo.Create(message)
}

func (uc *MessageUseCase) GetMessages(receiverID int) ([]entities.Message, error) {
	return uc.messageRepo.GetByReceiverID(receiverID)
}
