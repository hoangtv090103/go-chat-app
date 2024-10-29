//go:generate mockgen -source=chat_usecase.go -destination=./mocks/chat_store.go

package chatusecase

import (
	"context"
	"go-chat-app/internal/authentication/authinterfaces"
	chatdomain "go-chat-app/internal/chat/domain"
	chatinterfaces "go-chat-app/internal/chat/interfaces"
)

type IChatUseCase interface {
	SendMessage(context context.Context, message *chatdomain.ChatMessage) error
	GetMessagesByRoom(context context.Context, roomID uint) ([]chatdomain.ChatMessage, error)
	SendPrivateMessage(context context.Context, msg *chatdomain.ChatMessage) error
}
type ChatUseCase struct {
	chatRepo chatinterfaces.ChatMessageRepository
}

func NewChatUseCase(chatRepo chatinterfaces.ChatMessageRepository, userRepo authinterfaces.IUserRepository) *ChatUseCase {
	return &ChatUseCase{
		chatRepo: chatRepo,
	}
}

func (uc *ChatUseCase) SendMessage(context context.Context, message *chatdomain.ChatMessage) error {
	return uc.chatRepo.Store(context, message)
}

func (uc *ChatUseCase) GetMessagesByRoom(context context.Context, roomID uint) ([]chatdomain.ChatMessage, error) {
	return uc.chatRepo.GetMessagesByRoom(context, roomID)
}

func (uc *ChatUseCase) SendPrivateMessage(context context.Context, msg *chatdomain.ChatMessage) error {
	// Store the private message in the database
	return uc.chatRepo.Store(context, msg)
}
