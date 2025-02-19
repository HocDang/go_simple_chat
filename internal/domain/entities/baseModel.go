package entities

import (
	"time"

	"github.com/go-pg/pg/v10"
)

type BaseModel struct {
	CreatedAt time.Time `json:"created_at" pg:",default:now()"`
	UpdatedAt time.Time `json:"updated_at" pg:",default:now()"`
}

func (b *BaseModel) BeforeUpdate(db *pg.DB) error {
	b.UpdatedAt = time.Now()
	return nil
}
