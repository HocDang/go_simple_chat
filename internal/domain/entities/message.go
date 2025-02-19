package entities

import (
	"context"
	"errors"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

type Message struct {
	ID         uuid.UUID `json:"id" pg:",pk,type:uuid,default:gen_random_uuid()"`
	SenderID   uuid.UUID `json:"sender_id" pg:",notnull"`
	ReceiverID uuid.UUID `json:"receiver_id" pg:",notnull"`
	Content    string    `json:"content" pg:",notnull"`
	BaseModel

	Sender   *User `pg:"rel:has-one, fk:sender_id"`
	Receiver *User `pg:"rel:has-one, fk:receiver_id"`
}

func (m *Message) BeforeInsert(ctx context.Context, db *pg.DB) error {
	var sender, receiver User

	err := db.Model(&sender).Where("id = ?", m.SenderID).Select()
	if err != nil {
		return errors.New("invalid sender_id: user does not exist")
	}

	err = db.Model(&receiver).Where("id = ?", m.ReceiverID).Select()
	if err != nil {
		return errors.New("invalid receiver_id: user does not exist")
	}

	return nil
}
