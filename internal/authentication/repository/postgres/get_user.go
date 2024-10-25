package authrepository

import authdomain "go-chat-app/internal/authentication/domain"

func (r *AuthPGRepository) FindByUsername(username string) (*authdomain.User, error) {
	var user authdomain.User

	err := r.db.Where("username = ?", username).First(&user).Error

	return &user, err
}
