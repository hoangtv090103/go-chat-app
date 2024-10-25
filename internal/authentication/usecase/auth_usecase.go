package authusecase

import (
	"errors"
	authdomain "go-chat-app/internal/authentication/domain"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
)

type AuthUsecase struct {
	userRepo UserRepository
}

func NewAuthUseCase(userRepo UserRepository) *AuthUsecase {
	return &AuthUsecase{
		userRepo: userRepo,
	}
}

func (uc *AuthUsecase) Register(username, password string) error {
	// Check if the user already exists
	_, err := uc.userRepo.FindByUsername(username)
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
	return uc.userRepo.Store(user)
}

func (uc *AuthUsecase) Login(username, password string) (string, error) {
	// Find the user by username
	user, err := uc.userRepo.FindByUsername(username)

	if err != nil {
		return "", errors.New("user not exists")
	}

	// Compare the password and hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("incorrect password")
	}

	// Generate JWT Token
	// Define custom claims
	type CustomClaims struct {
		Username string `json:"username"`
		UserID   uint   `json:"user_id"`
		jwt.StandardClaims
	}

	// Create claims with custom fields
	claims := &CustomClaims{
		Username: username,
		UserID:   user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// TODO: What SigningMethodHS256 does?
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
