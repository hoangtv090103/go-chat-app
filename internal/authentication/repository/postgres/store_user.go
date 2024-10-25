package authrepository

import authdomain "go-chat-app/internal/authentication/domain"

func (r *AuthPGRepository) Store(user *authdomain.User) error {
	return r.db.Create(user).Error
}
