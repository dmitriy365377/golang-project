package service

import (
	"errors"
	"sync"
	"time"

	"golang-chat/internal/chat/model"

	"github.com/google/uuid"
)

type ChatService struct {
	chats    map[string]*model.Chat
	messages map[string]*model.Message
	mu       sync.RWMutex
}

func NewChatService() *ChatService {
	return &ChatService{
		chats:    make(map[string]*model.Chat),
		messages: make(map[string]*model.Message),
	}
}

func (s *ChatService) CreateChat(name, createdBy string, participants []string) (*model.Chat, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	chat := &model.Chat{
		ID:           uuid.New().String(),
		Name:         name,
		CreatedBy:    createdBy,
		CreatedAt:    time.Now(),
		Participants: append([]string{createdBy}, participants...),
	}

	s.chats[chat.ID] = chat
	return chat, nil
}

func (s *ChatService) ConnectChat(chatID, userID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	chat, exists := s.chats[chatID]
	if !exists {
		return errors.New("chat not found")
	}

	// Проверяем, является ли пользователь участником чата
	for _, participant := range chat.Participants {
		if participant == userID {
			return nil // Уже участник
		}
	}

	// Добавляем пользователя в участники
	chat.Participants = append(chat.Participants, userID)
	return nil
}

func (s *ChatService) SendMessage(chatID, userID, content string) (*model.Message, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	chat, exists := s.chats[chatID]
	if !exists {
		return nil, errors.New("chat not found")
	}

	// Проверяем, является ли пользователь участником чата
	isParticipant := false
	for _, participant := range chat.Participants {
		if participant == userID {
			isParticipant = true
			break
		}
	}

	if !isParticipant {
		return nil, errors.New("user is not a participant of this chat")
	}

	message := &model.Message{
		ID:        uuid.New().String(),
		ChatID:    chatID,
		UserID:    userID,
		Content:   content,
		CreatedAt: time.Now(),
	}

	s.messages[message.ID] = message
	return message, nil
}

func (s *ChatService) GetMessages(chatID string, limit, offset int) ([]*model.Message, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var messages []*model.Message
	for _, message := range s.messages {
		if message.ChatID == chatID {
			messages = append(messages, message)
		}
	}

	// Простая пагинация (в реальном приложении должна быть в базе данных)
	if offset >= len(messages) {
		return []*model.Message{}, nil
	}

	end := offset + limit
	if end > len(messages) {
		end = len(messages)
	}

	return messages[offset:end], nil
}
