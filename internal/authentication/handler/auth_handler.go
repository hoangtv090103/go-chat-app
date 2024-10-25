package authhandler

import (
	authdomain "go-chat-app/internal/authentication/domain"
	authusecase "go-chat-app/internal/authentication/usecase"

	"github.com/gofiber/fiber/v2"
)

func AuthHandler(app *fiber.App, authUsecase *authusecase.AuthUsecase) {
	// Register a new user
	app.Post("/register", func(c *fiber.Ctx) error {
		var data authdomain.UserCreate

		if err := c.BodyParser(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON, please check you request body.",
			})
		}

		err := authUsecase.Register(data.Username, data.Password)

		if err != nil {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "User registered successfully",
		})
	})

	// Login a user and return JWT token
	app.Post("/login", func(c *fiber.Ctx) error {
		var data authdomain.UserCreate

		if err := c.BodyParser(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}

		token, err := authUsecase.Login(data.Username, data.Password)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"token": token,
		})
	})
}
