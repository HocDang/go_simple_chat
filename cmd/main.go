package main

import (
	"chat-server/config"
	"chat-server/internal/app"
	"log"
)

func main() {

	// Load config from .env file
	env, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	config.PrintConfig(env)

	app.StartServer(env)
}
