package models

import "gorm.io/gorm"

// Migrate runs GORM AutoMigrate for all models in dependency order.
// Call this once on application startup.
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		// Core auth
		&User{},
		&Account{},
		&Session{},

		// Hubs
		&Hub{},

		// Staff
		&Admin{},

		// Reusable address
		&Address{},

		// Profiles (depend on User)
		&Merchant{},
		&Rider{},
		&GuestSender{},

		// Fraud
		&ReceiverFraudProfile{},

		// Shipment & related (depend on Merchant, Rider, GuestSender)
		&Shipment{},
		&DeliveryConfirmation{},
		&CodDeliveryConfirmation{},
		&WalkInPayment{},

		// Ratings (depend on Shipment, Rider, Merchant)
		&DeliveryRating{},
		&MerchantDeliveryRating{},

		// Finance
		&Withdrawal{},
	)
}
