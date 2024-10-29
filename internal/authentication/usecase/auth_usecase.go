package authrepository

import (
	"context"
	"errors"
	"go-chat-app/internal/authentication/authinterfaces"
	authdomain "go-chat-app/internal/authentication/domain"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	userRepo authinterfaces.IUserRepository
}

type IAuthUsecase interface {
	Register(context context.Context, username, password string) error
	Login(context context.Context, username, password string) (string, error)
}

func NewAuthUseCase(userRepo authinterfaces.IUserRepository) *AuthUsecase {
	return &AuthUsecase{
		userRepo: userRepo,
	}
}

func (uc *AuthUsecase) Register(context context.Context, username, password string) error {
	// Check if the user already exists
	_, err := uc.userRepo.FindByUsername(context, username)
	if err == nil {
		return errors.New("user already exists")
	}

	// Hash the pw
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create a new user
	user := &authdomain.User{
		Username: username,
		Password: string(hashedPassword),
	}

	// Store the user
	return uc.userRepo.Store(context, user)
}

func (uc *AuthUsecase) Login(context context.Context, username, password string) (string, error) {
	// Find the user by username
	user, err := uc.userRepo.FindByUsername(context, username)
	if err != nil {
		return "", errors.New("user not exists")
	}

	// Compare the password and hashed password
	if !uc.userRepo.CheckPasswordHash(context, password, user.Password) {
		return "", errors.New("incorrect password")
	}

	// Generate JWT Token
	type CustomClaims struct {
		Username string `json:"username"`
		UserID   uint   `json:"user_id"`
		jwt.StandardClaims
	}

	claims := &CustomClaims{
		Username: username,
		UserID:   user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
