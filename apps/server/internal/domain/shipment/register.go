package shipment

import (
	"server/internal/models"
	"server/pkg/middlewares"
	"server/pkg/utils"
	"server/pkg/validation"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func ShipmentRoutes(
	api fiber.Router,
	db *gorm.DB,
	v *validation.CustomValidator,
	jwt utils.JwtService,
) {
	repository := NewRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)

	router := api.Group("/shipments")

	// Public routes
	router.Post("/calculate-price", handler.CalculatePrice)
	router.Get("/track/:trackingCode", handler.TrackShipment)

	// Protected routes
	router.Use(middlewares.Authentication(jwt))
	router.Use(func(c fiber.Ctx) error {
		c.Locals("validator", v)
		return c.Next()
	})

	router.Post("/", handler.CreateShipment)
	router.Get("/", handler.ListShipments)
	router.Get("/:id", handler.GetShipmentByID)
	router.Put("/:id/status", handler.UpdateStatus)
	router.Put("/:id/assign-rider",
		middlewares.Authorization(string(models.RoleAdmin), string(models.RoleSuperAdmin)),
		handler.AssignRider,
	)
	router.Put("/:id/assign-hub",
		middlewares.Authorization(string(models.RoleAdmin), string(models.RoleSuperAdmin)),
		handler.AssignHub,
	)
}
