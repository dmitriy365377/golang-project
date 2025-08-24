package main

import (
	"log"

	"golang-chat/internal/client"
	"golang-chat/pkg/config"
)

func main() {
	cfg := config.Load()

	client := client.NewChatClient(cfg)

	if err := client.Run(); err != nil {
		log.Fatalf("Failed to run client: %v", err)
	}
}
