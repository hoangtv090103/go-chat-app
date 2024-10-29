//go:generate mockgen -source=room_repository.go -destination=./mocks/room_repository.go

package chatinterfaces

import (
	"context"
	chatdomain "go-chat-app/internal/chat/domain"
)

type RoomRepository interface {
	Store(context context.Context, room *chatdomain.RoomCreate) error
	FindByName(context context.Context, name string) (*chatdomain.Room, error)
	GetAll(context context.Context) ([]chatdomain.Room, error)
}
