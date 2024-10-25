package chatrepository

import chatdomain "go-chat-app/internal/chat/domain"

func (r *ChatPGRepository) Store(message *chatdomain.ChatMessage) error {
	return r.db.Create(message).Error
}
