package chatrepository

import (
	"context"
	chatdomain "go-chat-app/internal/chat/domain"
)

func (r *RoomPGRepository) Store(context context.Context, room *chatdomain.RoomCreate) error {
	return r.db.Table("rooms").Create(room).Error
}
