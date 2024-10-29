package chatrepository

import (
	"github.com/stretchr/testify/suite"
	mock_chatinterfaces "go-chat-app/internal/chat/interfaces/mocks"
	chatusecase "go-chat-app/internal/chat/usecase"
	"go.uber.org/mock/gomock"
	"testing"
)

type RoomSuite struct {
	suite.Suite
	mockRepo    *mock_chatinterfaces.MockRoomRepository
	roomUseCase chatusecase.IRoomUseCase
}

func TestRoomSuite(t *testing.T) {
	suite.Run(t, new(RoomSuite))
}

func (suite *RoomSuite) SetupSuite() {
	ctrl := gomock.NewController(suite.T())
	mockRepo := mock_chatinterfaces.NewMockRoomRepository(ctrl)
	usecase := chatusecase.NewRoomUseCase(mockRepo)
	suite.mockRepo = mockRepo
	suite.roomUseCase = usecase
}

func (suite *RoomSuite) Test_Create() {
	type inputArgs struct {
		roomName string
	}
}
