package authrepository

import "gorm.io/gorm"

type AuthPGRepository struct {
	db *gorm.DB
}

func NewAuthPGRepository(db *gorm.DB) *AuthPGRepository {
	return &AuthPGRepository{
		db: db,
	}
}
