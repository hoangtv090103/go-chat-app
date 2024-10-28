package authrepository

import (
	"context"
	authdomain "go-chat-app/internal/authentication/domain"
)

func (r *AuthPGRepository) Store(context context.Context, user *authdomain.User) error {
	return r.db.Create(user).Error
}
