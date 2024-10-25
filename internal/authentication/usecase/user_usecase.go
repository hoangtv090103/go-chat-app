package authusecase

import (
	authdomain "go-chat-app/internal/authentication/domain"
)

type UserRepository interface {
	Store(user *authdomain.User) error
	FindByUsername(username string) (*authdomain.User, error)
}

// type UserUseCase struct {
// 	useRepo UserRepository
// }
