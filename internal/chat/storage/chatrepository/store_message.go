package chatrepository

import (
	"context"
	chatdomain "go-chat-app/internal/chat/domain"
)

func (r *ChatPGRepository) Store(context context.Context, message *chatdomain.ChatMessage) error {
	return r.db.Create(message).Error
}
