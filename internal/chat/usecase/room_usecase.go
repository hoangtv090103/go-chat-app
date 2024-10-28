//go:generate mockgen -source=room_usecase.go -destination=./mocks/room_store.go

package chatusecase

import (
	"errors"
	chatdomain "go-chat-app/internal/chat/domain"
)

type RoomRepository interface {
	Store(room *chatdomain.RoomCreate) error
	FindByName(name string) (*chatdomain.Room, error)
	GetAll() ([]chatdomain.Room, error)
}

type RoomUseCase struct {
	roomRepo RoomRepository
}

func NewRoomUseCase(roomRepo RoomRepository) *RoomUseCase {
	return &RoomUseCase{
		roomRepo: roomRepo,
	}
}

func (uc *RoomUseCase) CreateRoom(roomName string) error {
	_, err := uc.roomRepo.FindByName(roomName)

	if err == nil {
		return errors.New("room already exists")
	}

	room := &chatdomain.RoomCreate{
		Name: roomName,
	}

	return uc.roomRepo.Store(room)
}

func (uc *RoomUseCase) GetAllRooms() ([]chatdomain.Room, error) {
	return uc.roomRepo.GetAll()
}

func (uc *RoomUseCase) GetRoomByName(roomName string) (*chatdomain.Room, error) {
	return uc.roomRepo.FindByName(roomName)
}
