package searches

import (
	"chat-server/internal/domain/entities"

	"github.com/google/uuid"
)

type MessageSearch interface {
	IndexMessage(msg entities.Message) error
	SearchMessages(senderID uuid.UUID, receiverID uuid.UUID, keyword string) ([]entities.Message, error)
}
