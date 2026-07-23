package rating

import (
	"server/internal/models"
	"server/pkg/middlewares"
	"server/pkg/utils"
	"server/pkg/validation"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func RatingRoutes(
	api fiber.Router,
	db *gorm.DB,
	v *validation.CustomValidator,
	jwt utils.JwtService,
) {
	repository := NewRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)

	router := api.Group("/ratings")

	// Public rating route for receivers
	router.Post("/rider", handler.RateRider)

	router.Use(middlewares.Authentication(jwt))
	router.Use(func(c fiber.Ctx) error {
		c.Locals("validator", v)
		return c.Next()
	})

	router.Get("/rider/:riderId", handler.ListRiderRatings)
	router.Post("/merchant", handler.RatePlatform)
	router.Get("/merchant",
		middlewares.Authorization(string(models.RoleAdmin), string(models.RoleSuperAdmin)),
		handler.ListMerchantRatings,
	)
}
