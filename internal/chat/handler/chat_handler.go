package handler

import (
	"context"

	"golang-chat/internal/chat/service"
	"golang-chat/proto/chat"
)

type ChatHandler struct {
	chat.UnimplementedChatServiceServer

	chatService *service.ChatService
}

func NewChatHandler(chatService *service.ChatService) *ChatHandler {
	return &ChatHandler{
		chatService: chatService,
	}
}

func (h *ChatHandler) CreateChat(ctx context.Context, req *chat.CreateChatRequest) (*chat.CreateChatResponse, error) {
	chatModel, err := h.chatService.CreateChat(req.Name, req.CreatedBy, req.Participants)
	if err != nil {
		return &chat.CreateChatResponse{Error: err.Error()}, nil
	}

	return &chat.CreateChatResponse{
		Chat: &chat.Chat{
			Id:           chatModel.ID,
			Name:         chatModel.Name,
			CreatedBy:    chatModel.CreatedBy,
			CreatedAt:    chatModel.CreatedAt.Format("2006-01-02T15:04:05Z"),
			Participants: chatModel.Participants,
		},
	}, nil
}

func (h *ChatHandler) ConnectChat(ctx context.Context, req *chat.ConnectChatRequest) (*chat.ConnectChatResponse, error) {
	err := h.chatService.ConnectChat(req.ChatId, req.UserId)
	if err != nil {
		return &chat.ConnectChatResponse{Error: err.Error()}, nil
	}

	return &chat.ConnectChatResponse{
		Success: true,
	}, nil
}

func (h *ChatHandler) SendMessage(ctx context.Context, req *chat.SendMessageRequest) (*chat.SendMessageResponse, error) {
	message, err := h.chatService.SendMessage(req.ChatId, req.UserId, req.Content)
	if err != nil {
		return &chat.SendMessageResponse{Error: err.Error()}, nil
	}

	return &chat.SendMessageResponse{
		Message: &chat.Message{
			Id:        message.ID,
			ChatId:    message.ChatID,
			UserId:    message.UserID,
			Content:   message.Content,
			CreatedAt: message.CreatedAt.Format("2006-01-02T15:04:05Z"),
		},
	}, nil
}

func (h *ChatHandler) GetMessages(ctx context.Context, req *chat.GetMessagesRequest) (*chat.GetMessagesResponse, error) {
	messages, err := h.chatService.GetMessages(req.ChatId, int(req.Limit), int(req.Offset))
	if err != nil {
		return &chat.GetMessagesResponse{Error: err.Error()}, nil
	}

	var protoMessages []*chat.Message
	for _, msg := range messages {
		protoMessages = append(protoMessages, &chat.Message{
			Id:        msg.ID,
			ChatId:    msg.ChatID,
			UserId:    msg.UserID,
			Content:   msg.Content,
			CreatedAt: msg.CreatedAt.Format("2006-01-02T15:04:05Z"),
		})
	}

	return &chat.GetMessagesResponse{
		Messages: protoMessages,
	}, nil
}
