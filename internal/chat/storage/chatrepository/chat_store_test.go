package chatrepository

import (
	"context"
	"errors"
	"github.com/stretchr/testify/suite"
	mock_interfaces "go-chat-app/internal/authentication/authinterfaces/mocks"
	chatdomain "go-chat-app/internal/chat/domain"
	mock_chatinterfaces "go-chat-app/internal/chat/interfaces/mocks"
	chatusecase "go-chat-app/internal/chat/usecase"
	"go.uber.org/mock/gomock"
	"testing"
)

type ChatSuite struct {
	suite.Suite
	chatRepo    *mock_chatinterfaces.MockChatMessageRepository
	userRepo    *mock_interfaces.MockIUserRepository
	chatUsecase chatusecase.IChatUseCase
}

func TestChatSuite(t *testing.T) {
	suite.Run(t, new(ChatSuite))
}

func (suite *ChatSuite) SetupSuite() {
	ctrl := gomock.NewController(suite.T())
	userRepo := mock_interfaces.NewMockIUserRepository(ctrl)
	chatRepo := mock_chatinterfaces.NewMockChatMessageRepository(ctrl)
	usecase := chatusecase.NewChatUseCase(chatRepo, userRepo)
	suite.chatRepo = chatRepo
	suite.userRepo = userRepo
	suite.chatUsecase = usecase
}

func (suite *ChatSuite) Test_SendMessages() {
	ctx := context.Background()
	type inputArgs struct {
		message *chatdomain.ChatMessage
	}

	tests := []struct {
		name        string
		args        inputArgs
		mockSetup   func()
		expectedErr error
	}{
		{
			name: "Successful message send",
			args: inputArgs{
				message: &chatdomain.ChatMessage{
					SenderID:  1,
					RoomID:    1,
					Message:   "Good morning",
					IsPrivate: false,
				},
			},
			mockSetup: func() {
				suite.chatRepo.EXPECT().Store(ctx, gomock.Any()).Return(nil)
			},
		},
		{
			name: "Failed message send",
			args: inputArgs{
				message: &chatdomain.ChatMessage{
					SenderID:  1,
					RoomID:    1,
					Message:   "Goodbye",
					IsPrivate: false,
				},
			},
			mockSetup: func() {
				suite.chatRepo.EXPECT().Store(ctx, gomock.Any()).Return(errors.New("failed to store message"))
			},
			expectedErr: errors.New("failed to store message"),
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name,
			func() {
				tt.mockSetup()
				err := suite.chatUsecase.SendMessage(ctx, tt.args.message)
				suite.Equal(tt.expectedErr, err)
			})
	}
}

func (suite *ChatSuite) Test_GetMessagesByRoom() {
	ctx := context.Background()
	type inputArgs struct {
		roomID uint
	}

	tests := []struct {
		name             string
		args             inputArgs
		mockSetup        func()
		expectedErr      error
		expectedMessages []chatdomain.ChatMessage
	}{
		{
			name: "Successful get messages",
			args: inputArgs{
				roomID: 1,
			},
			mockSetup: func() {
				messages := []chatdomain.ChatMessage{
					{SenderID: 1, RoomID: 1, Message: "Hi"},
					{SenderID: 2, RoomID: 1, Message: "Bye"},
				}

				suite.chatRepo.EXPECT().GetMessagesByRoom(ctx, uint(1)).Return(messages, nil)
			},
			expectedErr: nil,
			expectedMessages: []chatdomain.ChatMessage{
				{SenderID: 1, RoomID: 1, Message: "Hi"},
				{SenderID: 2, RoomID: 1, Message: "Bye"},
			},
		},
		{
			name: "Failed get messages",
			args: inputArgs{
				roomID: 1,
			},
			mockSetup: func() {
				suite.chatRepo.EXPECT().GetMessagesByRoom(ctx, uint(1)).Return(nil, errors.New("failed to get messages"))
			},
			expectedErr:      errors.New("failed to get messages"),
			expectedMessages: nil,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			tt.mockSetup()
			messages, err := suite.chatUsecase.GetMessagesByRoom(ctx, tt.args.roomID)
			suite.Equal(tt.expectedErr, err)
			suite.Equal(tt.expectedMessages, messages)
		})
	}
}

func (suite *ChatSuite) Test_SendPrivateMessage() {
	type inputArgs struct {
	}
}
