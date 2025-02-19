package postgres

import (
	"chat-server/internal/domain/entities"
	"chat-server/internal/domain/repositories"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

type messagePgRepository struct {
	db *pg.DB
}

func NewMessagePgRepository(db *pg.DB) repositories.MessageRepository {
	return &messagePgRepository{db: db}
}

func (r *messagePgRepository) Create(message *entities.Message) error {
	_, err := r.db.Model(message).Insert()
	return err
}

func (r *messagePgRepository) GetByReceiverID(receiverID uuid.UUID) ([]entities.Message, error) {
	var messages []entities.Message
	err := r.db.Model(&messages).Where("receiver_id = ?", receiverID).Select()
	if err != nil {
		return nil, err
	}
	return messages, nil
}
