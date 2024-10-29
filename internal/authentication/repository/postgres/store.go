package authrepository

import (
	"go-chat-app/internal/authentication/authinterfaces"
	"gorm.io/gorm"
)

type AuthPGRepository struct {
	db *gorm.DB
}

func NewAuthPGRepository(db *gorm.DB) authinterfaces.IUserRepository {
	return &AuthPGRepository{
		db: db,
	}
}
