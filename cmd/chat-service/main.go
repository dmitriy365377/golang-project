package main

import (
	"log"
	"net"

	"golang-chat/internal/chat/handler"
	"golang-chat/internal/chat/service"
	"golang-chat/pkg/config"
	"golang-chat/proto/chat"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	lis, err := net.Listen("tcp", cfg.ChatServicePort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	chatService := service.NewChatService()
	chatHandler := handler.NewChatHandler(chatService)

	chat.RegisterChatServiceServer(grpcServer, chatHandler)

	log.Printf("Chat Service starting on %s", cfg.ChatServicePort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
