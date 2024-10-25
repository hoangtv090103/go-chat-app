package main

import (
	dbconfig "go-chat-app/configs/db"
	authhandler "go-chat-app/internal/authentication/handler"
	authrepository "go-chat-app/internal/authentication/repository/postgres"
	authusecase "go-chat-app/internal/authentication/usecase"
	chathandler "go-chat-app/internal/chat/handler"
	"go-chat-app/internal/chat/storage/chatrepository"
	chatusecase "go-chat-app/internal/chat/usecase"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

const port = ":3001"

func main() {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatal("Cannot load environment variable", err)
	}

	// Initialize PostgresDB
	db := dbconfig.InitPostgresDB()

	// Initialize repository
	userRepo := authrepository.NewAuthPGRepository(db)
	chatRepo := chatrepository.NewChatPGRepository(db)
	roomRepo := chatrepository.NewRoomPGRepository(db)

	// Initialize use cases
	authUseCase := authusecase.NewAuthUseCase(userRepo)
	chatUseCase := chatusecase.NewChatUseCase(chatRepo, userRepo)
	roomUseCase := chatusecase.NewRoomUseCase(roomRepo)

	// Create Fiber app
	app := fiber.New()
	// Enable CORS with default options
	app.Use(cors.New())

	// Register Authentication routes
	authhandler.AuthHandler(app, authUseCase)

	// Register Websocket routes
	chathandler.WebSocketHandler(app, chatUseCase)

	// Register Room routes
	chathandler.RoomHandler(app, roomUseCase)

	app.Listen(port)
}
