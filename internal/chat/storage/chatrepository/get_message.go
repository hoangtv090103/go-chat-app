package chatrepository

import chatdomain "go-chat-app/internal/chat/domain"

func (r *ChatPGRepository) GetMessagesByRoom(roomID uint) ([]chatdomain.ChatMessage, error) {
	var messages []chatdomain.ChatMessage

	err := r.db.Where("room_id = ?", roomID).Find(&messages).Error

	return messages, err
}
