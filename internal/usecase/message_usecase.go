package usecase

import (
	"chat-server/internal/domain/entities"
	"chat-server/internal/domain/repositories"
	"chat-server/internal/domain/searches"

	"github.com/google/uuid"
)

type MessageUseCase struct {
	messageRepo   repositories.MessageRepository
	searchMessage searches.MessageSearch
}

func NewMessageUseCase(messageRepo repositories.MessageRepository, searchMessage searches.MessageSearch) *MessageUseCase {
	return &MessageUseCase{
		messageRepo:   messageRepo,
		searchMessage: searchMessage,
	}
}

func (uc *MessageUseCase) SendMessage(senderID uuid.UUID, receiverID uuid.UUID, content string) error {
	message := &entities.Message{
		SenderID:   senderID,
		ReceiverID: receiverID,
		Content:    content,
	}

	ms, err := uc.messageRepo.Create(message)
	if err != nil {
		return err
	}

	err = uc.searchMessage.IndexMessage(*ms)
	if err != nil {
		return err
	}

	return nil
}

func (uc *MessageUseCase) GetMessage(receiverID uuid.UUID, senderID uuid.UUID) ([]entities.Message, error) {
	return uc.messageRepo.GetByReceiverID(receiverID, senderID)
}

func (uc *MessageUseCase) SearchMessage(receiverID uuid.UUID, senderID uuid.UUID, query string) ([]entities.Message, error) {
	return uc.searchMessage.SearchMessages(receiverID, senderID, query)
}
