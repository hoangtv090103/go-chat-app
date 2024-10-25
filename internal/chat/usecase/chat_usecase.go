package chatusecase

import (
	userusecase "go-chat-app/internal/authentication/usecase"
	chatdomain "go-chat-app/internal/chat/domain"
)

type ChatMessageRepository interface {
	Store(message *chatdomain.ChatMessage) error
	GetMessagesByRoom(roomID uint) ([]chatdomain.ChatMessage, error)
}

type ChatUseCase struct {
	chatRepo ChatMessageRepository
}

func NewChatUseCase(chatRepo ChatMessageRepository, userRepo userusecase.UserRepository) *ChatUseCase {
	return &ChatUseCase{
		chatRepo: chatRepo,
	}
}

func (uc *ChatUseCase) SendMessage(message *chatdomain.ChatMessage) error {
	return uc.chatRepo.Store(message)
}

func (uc *ChatUseCase) GetMessagesByRoom(roomID uint) ([]chatdomain.ChatMessage, error) {
	return uc.chatRepo.GetMessagesByRoom(roomID)
}

func (uc *ChatUseCase) SendPrivateMessage(msg *chatdomain.ChatMessage) error {
	// Store the private message in the database
	return uc.chatRepo.Store(msg)
}
