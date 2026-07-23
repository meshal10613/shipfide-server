package admin

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"server/internal/config"
	"server/internal/models"
	"server/pkg/email"
	"server/pkg/middlewares"
	"server/pkg/utils"
	"server/pkg/validation"
)

func AdminRoutes(
	api fiber.Router,
	db *gorm.DB,
	v *validation.CustomValidator,
	jwt utils.JwtService,
) {
	repository := NewRepository(db)
	emailSender := email.NewEmailSender(config.AppConfig)
	service := NewService(repository, emailSender, config.AppConfig)
	handler := NewHandler(service)

	router := api.Group("/admins")

	// Apply Authentication middleware to all admin routes
	router.Use(middlewares.Authentication(jwt))

	// Middleware to inject validator into context
	router.Use(func(c fiber.Ctx) error {
		c.Locals("validator", v)
		return c.Next()
	})

	// Get all admins (SuperAdmin only)
	router.Get("/",
		middlewares.Authorization(string(models.RoleSuperAdmin)),
		handler.GetAdmins,
	)

	// Get admin by ID (SuperAdmin only)
	router.Get("/:id",
		middlewares.Authorization(string(models.RoleSuperAdmin)),
		handler.GetAdminByID,
	)

	// Create admin (SuperAdmin only)
	router.Post("/",
		middlewares.Authorization(string(models.RoleSuperAdmin)),
		handler.CreateAdmin,
	)

	// Update admin (SuperAdmin only)
	router.Put("/:id",
		middlewares.Authorization(string(models.RoleSuperAdmin)),
		handler.UpdateAdmin,
	)

	// Delete admin (SuperAdmin only)
	router.Delete("/:id",
		middlewares.Authorization(string(models.RoleSuperAdmin)),
		handler.DeleteAdmin,
	)
}
