package chatrepository

import (
	"context"
	chatdomain "go-chat-app/internal/chat/domain"
)

func (r *RoomPGRepository) FindByName(context context.Context, name string) (*chatdomain.Room, error) {
	var room chatdomain.Room
	err := r.db.Where("name = ?", name).First(&room).Error

	return &room, err
}

func (r *RoomPGRepository) GetAll(context context.Context) ([]chatdomain.Room, error) {
	var rooms []chatdomain.Room

	err := r.db.Find(&rooms).Error

	return rooms, err
}
