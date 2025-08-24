package main

import (
	"log"
	"net"

	"golang-chat/internal/auth/handler"
	"golang-chat/internal/auth/service"
	"golang-chat/pkg/config"
	"golang-chat/proto/auth"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	lis, err := net.Listen("tcp", cfg.AuthServicePort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	authService := service.NewAuthService()
	authHandler := handler.NewAuthHandler(authService)

	auth.RegisterAuthServiceServer(grpcServer, authHandler)
	auth.RegisterUserServiceServer(grpcServer, authHandler)
	auth.RegisterAccessServiceServer(grpcServer, authHandler)

	log.Printf("Auth Service starting on %s", cfg.AuthServicePort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
