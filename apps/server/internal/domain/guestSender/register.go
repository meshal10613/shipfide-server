package guestSender

import (
	"server/internal/models"
	"server/pkg/middlewares"
	"server/pkg/utils"
	"server/pkg/validation"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func GuestSenderRoutes(
	api fiber.Router,
	db *gorm.DB,
	v *validation.CustomValidator,
	jwt utils.JwtService,
) {
	repository := NewRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)

	router := api.Group("/guest-senders")

	router.Use(middlewares.Authentication(jwt))
	router.Use(middlewares.Authorization(string(models.RoleAdmin), string(models.RoleSuperAdmin)))
	router.Use(func(c fiber.Ctx) error {
		c.Locals("validator", v)
		return c.Next()
	})

	router.Post("/", handler.CreateGuestSender)
	router.Get("/", handler.ListGuestSenders)
	router.Get("/:id", handler.GetGuestSenderByID)
	router.Put("/:id/flag", handler.FlagGuestSender)
}
