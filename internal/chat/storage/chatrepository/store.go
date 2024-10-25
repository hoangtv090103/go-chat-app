package chatrepository

import "gorm.io/gorm"

type ChatPGRepository struct {
	db *gorm.DB
}

func NewChatPGRepository(db *gorm.DB) *ChatPGRepository {
	return &ChatPGRepository{
		db: db,
	}
}

type RoomPGRepository struct {
	db *gorm.DB
}

func NewRoomPGRepository(db *gorm.DB) *RoomPGRepository {
	return &RoomPGRepository{
		db: db,
	}
}
