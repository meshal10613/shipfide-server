package walkInPayment

import (
	"server/internal/models"
	"server/pkg/middlewares"
	"server/pkg/utils"
	"server/pkg/validation"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func WalkInPaymentRoutes(
	api fiber.Router,
	db *gorm.DB,
	v *validation.CustomValidator,
	jwt utils.JwtService,
) {
	repository := NewRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)

	router := api.Group("/payments/walk-in")

	router.Use(middlewares.Authentication(jwt))
	router.Use(middlewares.Authorization(string(models.RoleAdmin), string(models.RoleSuperAdmin)))
	router.Use(func(c fiber.Ctx) error {
		c.Locals("validator", v)
		return c.Next()
	})

	router.Post("/", handler.CreatePayment)
	router.Get("/", handler.ListPayments)
}
