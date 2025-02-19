package main

import (
	"chat-server/config"
	"chat-server/internal/app"
)

func main() {
	env := config.LoadConfig()

	app.StartServer(env)
}
