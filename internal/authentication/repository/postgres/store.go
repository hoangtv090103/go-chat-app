// internal/authentication/repository/postgres/store.go
package authrepository

import (
	"go-chat-app/internal/authentication/interfaces"
	"gorm.io/gorm"
)

type AuthPGRepository struct {
	db *gorm.DB
}

func NewAuthPGRepository(db *gorm.DB) interfaces.IUserRepository {
	return &AuthPGRepository{
		db: db,
	}
}
