package authrepository

import (
	"context"
	authdomain "go-chat-app/internal/authentication/domain"
	"golang.org/x/crypto/bcrypt"
)

func (r *AuthPGRepository) FindByUsername(context context.Context, username string) (*authdomain.User, error) {
	var user authdomain.User

	err := r.db.Where("username = ?", username).First(&user).Error

	return &user, err
}

func (r *AuthPGRepository) CheckPasswordHash(context context.Context, password, hash string) bool {
	// Compare the password and hashed password
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
