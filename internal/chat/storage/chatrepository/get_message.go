package chatrepository

import (
	"context"
	"errors"
	chatdomain "go-chat-app/internal/chat/domain"
)

func (r *ChatPGRepository) GetMessagesByRoom(context context.Context, roomID uint) ([]chatdomain.ChatMessage, error) {
	var messages []chatdomain.ChatMessage

	err := r.db.Where("room_id = ?", roomID).Find(&messages).Error
	if err != nil {
		return nil, errors.New("failed to get messages")
	}

	return messages, nil
}
