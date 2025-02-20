package migrations

import (
	"chat-server/internal/domain/entities"
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func MigrateTableMessages(db *pg.DB) error {
	model := (*entities.Message)(nil)
	err := db.Model(model).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})
	if err != nil {
		log.Println("❌ Error creating table messages: ", err)
	}

	log.Println("✅ Migration completed table messages")
	return nil
}
