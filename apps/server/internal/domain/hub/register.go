package hub

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"server/internal/models"
	"server/pkg/middlewares"
	"server/pkg/utils"
	"server/pkg/validation"
)

func HubRoutes(
	api fiber.Router,
	db *gorm.DB,
	v *validation.CustomValidator,
	jwt utils.JwtService,
) {
	repository := NewRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)

	router := api.Group("/hubs")

	// Apply Authentication middleware to all hub routes
	router.Use(middlewares.Authentication(jwt))

	// Get all hubs (Admin/SuperAdmin only)
	router.Get("/",
		middlewares.Authorization(string(models.RoleSuperAdmin)),
		handler.GetHubs,
	)

	// Get hub by ID (Admin/SuperAdmin only)
	router.Get("/:id",
		middlewares.Authorization(string(models.RoleSuperAdmin)),
		handler.GetHubByID,
	)

	// Create hub (Super Admin only)
	router.Post("/",
		middlewares.Authorization(string(models.RoleSuperAdmin)),
		handler.CreateHub,
	)

	// Update hub (Super Admin only)
	router.Put("/:id",
		middlewares.Authorization(string(models.RoleSuperAdmin)),
		handler.UpdateHub,
	)

	// Delete hub (Super Admin only)
	router.Delete("/:id",
		middlewares.Authorization(string(models.RoleSuperAdmin)),
		handler.DeleteHub,
	)
}
