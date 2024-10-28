//go:generate mockgen -source=user_repository.go -destination=./mocks/user_repository.go
package interfaces

import (
	"context"
	authdomain "go-chat-app/internal/authentication/domain"
)

type IUserRepository interface {
	Store(context context.Context, user *authdomain.User) error
	FindByUsername(context context.Context, username string) (*authdomain.User, error)
	CheckPasswordHash(context context.Context, password, hash string) bool
}
