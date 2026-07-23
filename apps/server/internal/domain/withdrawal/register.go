package withdrawal

import (
	"server/internal/models"
	"server/pkg/middlewares"
	"server/pkg/utils"
	"server/pkg/validation"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func WithdrawalRoutes(
	api fiber.Router,
	db *gorm.DB,
	v *validation.CustomValidator,
	jwt utils.JwtService,
) {
	repository := NewRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)

	router := api.Group("/withdrawals")

	router.Use(middlewares.Authentication(jwt))
	router.Use(func(c fiber.Ctx) error {
		c.Locals("validator", v)
		return c.Next()
	})

	router.Post("/", handler.CreateWithdrawal)
	router.Get("/", handler.ListWithdrawals)
	router.Get("/:id", handler.GetWithdrawalByID)
	router.Put("/:id/status",
		middlewares.Authorization(string(models.RoleAdmin), string(models.RoleSuperAdmin)),
		handler.UpdateStatus,
	)
}
