package user

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"server/internal/models"
	"server/pkg/middlewares"
	"server/pkg/utils"
)

func UserRoutes(
	api fiber.Router,
	db *gorm.DB,
	jwt utils.JwtService,
) {
	repository := NewRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)

	router := api.Group("/users")

	// Apply Authentication middleware to all user routes
	router.Use(middlewares.Authentication(jwt))

	// Get all users (Admin/SuperAdmin only)
	router.Get("/",
		middlewares.Authorization(string(models.RoleSuperAdmin)),
		handler.GetUsers,
	)

	// Get user by ID (Authenticated users)
	router.Get("/:id", handler.GetUserByID)

	// Update user (Self, Admin, or SuperAdmin)
	router.Put("/:id", handler.UpdateUser)

	// Delete user (Self, Admin, or SuperAdmin)
	router.Delete("/:id", handler.DeleteUser)
}
