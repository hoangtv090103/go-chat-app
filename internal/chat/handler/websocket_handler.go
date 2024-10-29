package chathandler

import (
	"context"
	"log"
	"strconv"

	chatdomain "go-chat-app/internal/chat/domain"
	chatusecase "go-chat-app/internal/chat/usecase"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

var clients = make(map[*websocket.Conn]map[uint]bool) // Track clients by room

func WebSocketHandler(app *fiber.App, chatUseCase *chatusecase.ChatUseCase) {
	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		tokenString := c.Query("token")
		roomID, err := strconv.Atoi(c.Query("room-id"))
		if err != nil {
			log.Println("Invalid room-id:", err)
			c.Close()
			return
		}

		claims := &struct {
			Username string `json:"username"`
			UserID   uint   `json:"user_id"`
			jwt.StandardClaims
		}{}
		_, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})
		if err != nil || claims.Valid() != nil {
			log.Println("Invalid JWT token:", err)
			c.Close()
			return
		}

		userID := claims.UserID // Extract user ID from JWT claims

		if _, ok := clients[c]; !ok {
			clients[c] = make(map[uint]bool)
		}
		clients[c][uint(roomID)] = true

		// Send message history to the client
		sendMessageHistory(c, chatUseCase, uint(roomID))

		// Handle incoming messages
		handleMessages(c, chatUseCase, uint(roomID), uint(userID))

		// Clean up on disconnect
		defer func() {
			delete(clients, c)
			c.Close()
		}()
	}))
}

// Function to send message history
func sendMessageHistory(conn *websocket.Conn, chatUseCase *chatusecase.ChatUseCase, roomID uint) {
	messages, err := chatUseCase.GetMessagesByRoom(context.Background(), roomID) // Pass context
	if err != nil {
		log.Println("Error fetching message history: ", err)
		return
	}

	for _, msg := range messages {
		err := conn.WriteJSON(msg)
		if err != nil {
			log.Printf("Error sending message history: %v", err)
			conn.Close()
			delete(clients, conn)
			return
		}
	}
}

func handleMessages(conn *websocket.Conn, chatUseCase *chatusecase.ChatUseCase, roomID uint, userID uint) {
	for {
		var msg chatdomain.ChatMessage

		err := conn.ReadJSON(&msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket closed unexpectedly: %v", err)
			} else {
				log.Printf("Error reading JSON: %v", err)
			}
			// Safely close the connection
			delete(clients, conn)
			conn.Close()
			break
		}

		// Set sender ID and room ID for the message
		msg.SenderID = userID
		msg.RoomID = roomID

		// Handle Room-based messages
		if !msg.IsPrivate {
			broadcastRoomMessage(conn, msg, roomID)
		} else {
			// Handle Direct Message (DM)
			sendDirectMessage(conn, msg, msg.RecipientID)
		}

		// Optionally, store the message in the database
		err = chatUseCase.SendMessage(context.Background(), &msg)

		if err != nil {
			log.Println("Error storing message: ", err)
		}
	}
}

func broadcastRoomMessage(_ *websocket.Conn, msg chatdomain.ChatMessage, roomID uint) {
	// Broadcast the message to all users in the room
	for client := range clients {
		if clients[client][roomID] {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Broadcast error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func sendDirectMessage(_ *websocket.Conn, msg chatdomain.ChatMessage, recipientID uint) {
	// Send the message only to the specific recipient
	for client := range clients {
		if clients[client][recipientID] { // Assuming clients are mapped by user ID for DMs
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("DM error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
