package session

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"server/pkg/middlewares"
	"server/pkg/utils"
)

func SessionRoutes(
	api fiber.Router,
	db *gorm.DB,
	jwt utils.JwtService,
) {
	repository := NewRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)

	router := api.Group("/sessions")

	// Apply Authentication middleware to all session routes
	router.Use(middlewares.Authentication(jwt))

	// Get all active sessions for the current user
	router.Get("/", handler.GetSessions)

	// Delete a specific session (own sessions only)
	router.Delete("/:id", handler.DeleteSession)
}
