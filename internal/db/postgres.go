package db

import (
	"chat-server/config"
	"log"

	"github.com/go-pg/pg/v10"
)

func InitPostgres(env *config.Env) *pg.DB {
	dbHost := env.PostgresHost
	dbPort := env.PostgresPort
	dbName := env.PostgresName
	dbUser := env.PostgresUser
	dbPass := env.PostgresPass

	db := pg.Connect(&pg.Options{
		Addr:     dbHost + ":" + dbPort,
		User:     dbUser,
		Password: dbPass,
		Database: dbName,
	})

	// Check connection
	_, err := db.Exec("SELECT 1")
	if err != nil {
		log.Println("❌ Postgres connection failed:", err)
		panic(err)
	}

	log.Println("✅ Connected to Postgres")
	return db
}

func Close(db *pg.DB) {
	db.Close()
}
