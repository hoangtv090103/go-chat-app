package chatrepository

import (
	chatdomain "go-chat-app/internal/chat/domain"
)

func (r *RoomPGRepository) Store(room *chatdomain.RoomCreate) error {
	return r.db.Table("rooms").Create(room).Error
}