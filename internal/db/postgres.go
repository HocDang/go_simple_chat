package db

import (
	"log"

	"github.com/go-pg/pg/v10"
)

var DB *pg.DB

func InitPostgres() {
	DB = pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "user",
		Password: "password",
		Database: "dbname",
	})

	if _, err := DB.Exec("SELECT 1"); err != nil {
		log.Fatal("❌ PostgreSQL connection failed:", err)
	}
	log.Println("✅ Connected to PostgreSQL")
}
