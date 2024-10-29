//go:generate mockgen -source=chat_repository.go -destination=./mocks/chat_repository.go

package chatinterfaces

import (
	"context"
	chatdomain "go-chat-app/internal/chat/domain"
)

type ChatMessageRepository interface {
	Store(context context.Context, message *chatdomain.ChatMessage) error
	GetMessagesByRoom(context context.Context, roomID uint) ([]chatdomain.ChatMessage, error)
}
