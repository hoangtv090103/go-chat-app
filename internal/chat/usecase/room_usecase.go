//go:generate mockgen -source=room_usecase.go -destination=./mocks/room_store.go

package chatusecase

import (
	"context"
	"errors"
	chatdomain "go-chat-app/internal/chat/domain"
	chatinterfaces "go-chat-app/internal/chat/interfaces"
)

type IRoomUseCase interface {
	CreateRoom(context context.Context, roomName string) error
	GetAllRooms(context context.Context) ([]chatdomain.Room, error)
	GetRoomByName(context context.Context, roomName string) (*chatdomain.Room, error)
}

type RoomUseCase struct {
	roomRepo chatinterfaces.RoomRepository
}

func NewRoomUseCase(roomRepo chatinterfaces.RoomRepository) *RoomUseCase {
	return &RoomUseCase{
		roomRepo: roomRepo,
	}
}

func (uc *RoomUseCase) CreateRoom(context context.Context, roomName string) error {
	_, err := uc.roomRepo.FindByName(context, roomName)

	if err == nil {
		return errors.New("room already exists")
	}

	room := &chatdomain.RoomCreate{
		Name: roomName,
	}

	return uc.roomRepo.Store(context, room)
}

func (uc *RoomUseCase) GetAllRooms(context context.Context) ([]chatdomain.Room, error) {
	return uc.roomRepo.GetAll(context)
}

func (uc *RoomUseCase) GetRoomByName(context context.Context, roomName string) (*chatdomain.Room, error) {
	return uc.roomRepo.FindByName(context, roomName)
}
