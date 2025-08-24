package client

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang-chat/pkg/config"
)

type ChatClient struct {
	config *config.Config
	reader *bufio.Reader
}

func NewChatClient(cfg *config.Config) *ChatClient {
	return &ChatClient{
		config: cfg,
		reader: bufio.NewReader(os.Stdin),
	}
}

func (c *ChatClient) Run() error {
	fmt.Println("Welcome to Golang Chat!")
	fmt.Println("Available commands:")
	fmt.Println("  login - Authenticate user")
	fmt.Println("  create chat <name> - Create new chat")
	fmt.Println("  connect <chat_id> - Connect to existing chat")
	fmt.Println("  send <chat_id> <message> - Send message to chat")
	fmt.Println("  quit - Exit application")

	for {
		fmt.Print("> ")
		input, err := c.reader.ReadString('\n')
		if err != nil {
			return err
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		if input == "quit" {
			fmt.Println("Goodbye!")
			break
		}

		if err := c.handleCommand(input); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}

	return nil
}

func (c *ChatClient) handleCommand(input string) error {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return nil
	}

	command := parts[0]

	switch command {
	case "login":
		return c.handleLogin(parts[1:])
	case "create":
		if len(parts) >= 3 && parts[1] == "chat" {
			return c.handleCreateChat(parts[2:])
		}
	case "connect":
		if len(parts) >= 2 {
			return c.handleConnect(parts[1])
		}
	case "send":
		if len(parts) >= 3 {
			return c.handleSendMessage(parts[1], strings.Join(parts[2:], " "))
		}
	default:
		fmt.Printf("Unknown command: %s\n", command)
	}

	return nil
}

func (c *ChatClient) handleLogin(args []string) error {
	if len(args) < 2 {
		fmt.Println("Usage: login <username> <password>")
		return nil
	}

	username := args[0]
	_ = args[1] // password - будет использоваться в gRPC вызове

	// Здесь должна быть логика аутентификации через gRPC
	fmt.Printf("Attempting to login as %s...\n", username)
	fmt.Println("(gRPC call to Auth Service would happen here)")

	return nil
}

func (c *ChatClient) handleCreateChat(args []string) error {
	if len(args) == 0 {
		fmt.Println("Usage: create chat <name>")
		return nil
	}

	name := strings.Join(args, " ")
	fmt.Printf("Creating chat: %s\n", name)
	fmt.Println("(gRPC call to Chat Service would happen here)")

	return nil
}

func (c *ChatClient) handleConnect(chatID string) error {
	fmt.Printf("Connecting to chat: %s\n", chatID)
	fmt.Println("(gRPC call to Chat Service would happen here)")

	return nil
}

func (c *ChatClient) handleSendMessage(chatID, message string) error {
	fmt.Printf("Sending message to chat %s: %s\n", chatID, message)
	fmt.Println("(gRPC call to Chat Service would happen here)")

	return nil
}
