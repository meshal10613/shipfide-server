package dto

import "server/internal/models"

type CreateDeliveryRatingRequest struct {
	ShipmentID    string  `json:"shipmentId" validate:"required"`
	RiderID       string  `json:"riderId" validate:"required"`
	ReceiverPhone string  `json:"receiverPhone" validate:"required"`
	IsAnonymous   bool    `json:"isAnonymous"`
	Stars         int     `json:"stars" validate:"required,min=1,max=5"`
	Comment       *string `json:"comment,omitempty"`
}

type CreateMerchantRatingRequest struct {
	ShipmentID string                    `json:"shipmentId" validate:"required"`
	MerchantID string                    `json:"merchantId" validate:"required"`
	Stars      int                       `json:"stars" validate:"required,min=1,max=5"`
	Comment    *string                   `json:"comment,omitempty"`
	Tag        *models.MerchantRatingTag `json:"tag,omitempty"`
}
