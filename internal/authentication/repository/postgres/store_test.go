package authrepository

import (
	"context"
	"errors"
	"github.com/stretchr/testify/suite"
	authdomain "go-chat-app/internal/authentication/domain"
	mock_interfaces "go-chat-app/internal/authentication/interfaces/mocks"
	authusecase "go-chat-app/internal/authentication/usecase"
	"go.uber.org/mock/gomock"
	testing "testing"
)

type UserSuite struct {
	suite.Suite
	mockRepo    *mock_interfaces.MockIUserRepository
	authUseCase authusecase.IAuthUsecase
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(UserSuite))
}

func (suite *UserSuite) SetupSuite() {
	ctrl := gomock.NewController(suite.T())
	mockRepo := mock_interfaces.NewMockIUserRepository(ctrl)
	usecase := authusecase.NewAuthUseCase(mockRepo)
	suite.mockRepo = mockRepo
	suite.authUseCase = usecase
}

func (suite *UserSuite) Test_Register() {
	type inputArgs struct {
		username string
		password string
	}

	tests := []struct {
		name        string
		args        inputArgs
		mockSetup   func()
		expectedErr error
	}{
		{
			name: "User already exists",
			args: inputArgs{
				username: "existinguser",
				password: "password",
			},
			mockSetup: func() {
				suite.mockRepo.EXPECT().FindByUsername(context.Background(), "existinguser").Return(&authdomain.User{}, nil)
			},
			expectedErr: errors.New("user already exists"),
		},
		{
			name: "Successful registration",
			args: inputArgs{
				username: "newuser",
				password: "password",
			},
			mockSetup: func() {
				suite.mockRepo.EXPECT().FindByUsername(context.Background(), "newuser").Return(nil, errors.New("not found"))
				suite.mockRepo.EXPECT().Store(context.Background(), gomock.Any()).Return(nil)
			},
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			tt.mockSetup()
			err := suite.authUseCase.Register(context.Background(), tt.args.username, tt.args.password)
			suite.Equal(tt.expectedErr, err)
		})
	}
}

// internal/authentication/repository/postgres/store_test.go
func (suite *UserSuite) Test_Login() {
	type inputArgs struct {
		username string
		password string
	}

	tests := []struct {
		name          string
		args          inputArgs
		mockSetup     func()
		expectedErr   error
		expectedToken string
	}{
		{
			name: "User not exists",
			args: inputArgs{
				username: "nonexistentuser",
				password: "password",
			},
			mockSetup: func() {
				suite.mockRepo.EXPECT().FindByUsername(context.Background(), "nonexistentuser").Return(nil, errors.New("user not exists"))
			},
			expectedErr:   errors.New("user not exists"),
			expectedToken: "",
		},
		{
			name: "Incorrect password",
			args: inputArgs{
				username: "existinguser",
				password: "wrongpassword",
			},
			mockSetup: func() {
				user := &authdomain.User{
					Username: "existinguser",
					Password: "$2a$10$7EqJtq98hPqEX7fNZaFWoOa1K8G1e1u1B1e1e1e1e1e1e1e1e1e1e", // bcrypt hash of "correctpassword"
				}
				suite.mockRepo.EXPECT().FindByUsername(context.Background(), "existinguser").Return(user, nil)
				suite.mockRepo.EXPECT().CheckPasswordHash(context.Background(), "wrongpassword", user.Password).Return(false)
			},
			expectedErr:   errors.New("incorrect password"),
			expectedToken: "",
		},
		{
			name: "Successful login",
			args: inputArgs{
				username: "existinguser",
				password: "correctpassword",
			},
			mockSetup: func() {
				user := &authdomain.User{
					Username: "existinguser",
					Password: "$2a$10$7EqJtq98hPqEX7fNZaFWoOa1K8G1e1u1B1e1e1e1e1e1e1e1e1e1e", // bcrypt hash of "correctpassword"
				}
				suite.mockRepo.EXPECT().FindByUsername(context.Background(), "existinguser").Return(user, nil)
				suite.mockRepo.EXPECT().CheckPasswordHash(context.Background(), "correctpassword", user.Password).Return(true)
			},
			expectedErr:   nil,
			expectedToken: ".eyJ1c2VybmFtZSI6ImV4aXN0aW5ndXNlciIsInBhc3N3b3JkIjoiJDJhJDEwJDdFcUp0cTk4aFBxRVg3Zk5aYUZXb09hMUs4RzFlMXUxQjFlMWUxZTFlMWUxZTFlMWUxZSIsImlhdCI6MTczMDA4NTE2N30.6V0kBoS4yGgVLur1MH3tztwzI47dDu5_Lmpuv910h3I", // Replace with actual token generation logic if needed
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			tt.mockSetup()
			token, err := suite.authUseCase.Login(context.Background(), tt.args.username, tt.args.password)
			suite.Equal(tt.expectedErr, err)
			if err == nil {
				suite.NotEmpty(token)
			} else {
				suite.Empty(token)
			}
		})
	}
}
