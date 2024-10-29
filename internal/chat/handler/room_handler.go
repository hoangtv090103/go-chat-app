package chathandler

import (
	chatusecase "go-chat-app/internal/chat/usecase"

	"github.com/gofiber/fiber/v2"
)

func RoomHandler(app *fiber.App, roomUseCase *chatusecase.RoomUseCase) {
	// Create a new chat room
	app.Post("/rooms", func(c *fiber.Ctx) error {
		data := new(struct {
			Name string `json:"name"`
		})

		if err := c.BodyParser(data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}

		err := roomUseCase.CreateRoom(c.Context(), data.Name)
		if err != nil {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		// Return the newly created room as response
		room, _ := roomUseCase.GetRoomByName(c.Context(), data.Name)
		return c.Status(fiber.StatusCreated).JSON(room)
	})

	// Get a list of all rooms
	app.Get("/rooms", func(c *fiber.Ctx) error {
		rooms, err := roomUseCase.GetAllRooms(c.Context())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Could not retrieve rooms",
			})
		}

		return c.Status(fiber.StatusOK).JSON(rooms)
	})
}
