package dbconfig

import (
	authdomain "go-chat-app/internal/authentication/domain"
	chatdomain "go-chat-app/internal/chat/domain"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgresDB() *gorm.DB {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(
		&authdomain.User{},
		&chatdomain.ChatMessage{},
		&chatdomain.Room{},
	)
	return db
}
