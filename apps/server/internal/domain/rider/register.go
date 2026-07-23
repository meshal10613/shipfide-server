package rider

import (
	"server/internal/models"
	"server/pkg/middlewares"
	"server/pkg/utils"
	"server/pkg/validation"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func RiderRoutes(
	api fiber.Router,
	db *gorm.DB,
	v *validation.CustomValidator,
	jwt utils.JwtService,
) {
	repository := NewRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)

	router := api.Group("/riders")

	router.Use(middlewares.Authentication(jwt))
	router.Use(func(c fiber.Ctx) error {
		c.Locals("validator", v)
		return c.Next()
	})

	router.Get("/me", handler.GetMyRiderProfile)
	router.Post("/", handler.CreateRider)

	router.Get("/",
		middlewares.Authorization(string(models.RoleAdmin), string(models.RoleSuperAdmin)),
		handler.ListRiders,
	)
	router.Get("/:id", handler.GetRiderByID)
	router.Put("/:id", handler.UpdateRider)
	router.Put("/:id/status", handler.UpdateStatus)
	router.Put("/:id/kyc",
		middlewares.Authorization(string(models.RoleAdmin), string(models.RoleSuperAdmin)),
		handler.UpdateKycStatus,
	)
}
