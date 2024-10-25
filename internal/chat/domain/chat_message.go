package chatdomain

import "time"

type ChatMessage struct {
	ID          uint      `gorm:"primaryKey"`
	SenderID    uint      `json:"sender_id" gorm:"not null"`
	RecipientID uint      `json:"recipient_id,omitempty"` // For private messaging
	RoomID      uint      `gorm:"not null"`
	Message     string    `json:"message" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	IsPrivate   bool      `json:"is_private"`
}
