package routes

import (
	"server/internal/config"
	"server/internal/domain/address"
	"server/internal/domain/admin"
	"server/internal/domain/auth"
	"server/internal/domain/deliveryConfirmation"
	"server/internal/domain/guestSender"
	"server/internal/domain/hub"
	"server/internal/domain/merchant"
	"server/internal/domain/rating"
	"server/internal/domain/receiverFraud"
	"server/internal/domain/rider"
	"server/internal/domain/session"
	"server/internal/domain/shipment"
	"server/internal/domain/upload"
	"server/internal/domain/user"
	"server/internal/domain/walkInPayment"
	"server/internal/domain/withdrawal"
	httpresponse "server/pkg/httpResponse"
	"server/pkg/utils"
	"server/pkg/validation"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

// SetupRoutes registers all application routes with required services
func SetupRoutes(
	app *fiber.App,
	db *gorm.DB,
	v *validation.CustomValidator,
	jwt utils.JwtService,
) {
	// Root endpoint
	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(httpresponse.Success{
			Success: true,
			Message: "Hello, World 👋!",
		})
	})

	// API base group
	api := app.Group("/api/v1")

	// Core & Identity routes
	auth.AuthRoutes(api, db, v, jwt)
	user.UserRoutes(api, db, jwt)
	session.SessionRoutes(api, db, jwt)
	hub.HubRoutes(api, db, v, jwt)
	admin.AdminRoutes(api, db, v, jwt)

	// Operational & Profile routes
	address.AddressRoutes(api, db, v, jwt)
	merchant.MerchantRoutes(api, db, v, jwt)
	rider.RiderRoutes(api, db, v, jwt)
	guestSender.GuestSenderRoutes(api, db, v, jwt)
	shipment.ShipmentRoutes(api, db, v, jwt)
	deliveryConfirmation.DeliveryConfirmationRoutes(api, db, v, jwt)
	walkInPayment.WalkInPaymentRoutes(api, db, v, jwt)
	receiverFraud.ReceiverFraudRoutes(api, db, v, jwt)
	rating.RatingRoutes(api, db, v, jwt)
	withdrawal.WithdrawalRoutes(api, db, v, jwt)

	// File & Image upload routes
	upload.UploadRoutes(api, config.AppConfig, jwt)
}
