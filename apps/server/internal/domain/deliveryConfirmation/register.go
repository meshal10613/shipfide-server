package deliveryConfirmation

import (
	"server/internal/models"
	"server/pkg/middlewares"
	"server/pkg/utils"
	"server/pkg/validation"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func DeliveryConfirmationRoutes(
	api fiber.Router,
	db *gorm.DB,
	v *validation.CustomValidator,
	jwt utils.JwtService,
) {
	repository := NewRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)

	router := api.Group("/delivery")

	router.Use(middlewares.Authentication(jwt))
	router.Use(func(c fiber.Ctx) error {
		c.Locals("validator", v)
		return c.Next()
	})

	router.Post("/otp/verify", handler.VerifyOtp)
	router.Post("/otp/regenerate",
		middlewares.Authorization(string(models.RoleAdmin), string(models.RoleSuperAdmin)),
		handler.RegenerateOtp,
	)

	router.Post("/cod/confirm", handler.ConfirmCod)
	router.Post("/cod/deposit", handler.DepositCod)
	router.Put("/cod/deposit/:id/approve",
		middlewares.Authorization(string(models.RoleAdmin), string(models.RoleSuperAdmin)),
		handler.ApproveDeposit,
	)
}
