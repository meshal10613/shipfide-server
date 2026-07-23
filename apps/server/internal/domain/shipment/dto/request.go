package dto

import (
	"server/internal/models"

	"github.com/shopspring/decimal"
)

type CalculatePriceRequest struct {
	Zone            models.ZoneType   `json:"zone" validate:"required"`
	WeightGrams     int               `json:"weightGrams" validate:"required,min=1"`
	ParcelType      models.ParcelType `json:"parcelType" validate:"required"`
	IsFragile       bool              `json:"isFragile"`
	PickupSurcharge *decimal.Decimal  `json:"pickupSurcharge,omitempty"`
}

type CreateShipmentRequest struct {
	SenderType    models.SenderType `json:"senderType" validate:"required"`
	MerchantID    *string           `json:"merchantId,omitempty"`
	GuestSenderID *string           `json:"guestSenderId,omitempty"`
	HubID         *string           `json:"hubId,omitempty"`

	ParcelType   models.ParcelType `json:"parcelType" validate:"required"`
	Weight       int               `json:"weight" validate:"required,min=1"` // grams
	Dimensions   *string           `json:"dimensions,omitempty"`
	ProductName  string            `json:"productName" validate:"required"`
	ProductValue *decimal.Decimal  `json:"productValue,omitempty"`
	IsFragile    bool              `json:"isFragile"`
	Notes        *string           `json:"notes,omitempty"`

	SenderName    string  `json:"senderName" validate:"required"`
	SenderPhone   string  `json:"senderPhone" validate:"required"`
	SenderAddress *string `json:"senderAddress,omitempty"`

	ReceiverName      string  `json:"receiverName" validate:"required"`
	ReceiverPhone     string  `json:"receiverPhone" validate:"required"`
	ReceiverEmail     *string `json:"receiverEmail,omitempty"`
	ReceiverAddress   string  `json:"receiverAddress" validate:"required"`
	ReceiverAddressID *string `json:"receiverAddressId,omitempty"`

	ZoneType  models.ZoneType `json:"zoneType" validate:"required"`
	CodAmount decimal.Decimal `json:"codAmount"`
}

type UpdateShipmentStatusRequest struct {
	Status models.ShipmentStatus `json:"status" validate:"required"`
	Notes  *string               `json:"notes,omitempty"`
}

type AssignRiderRequest struct {
	RiderID string `json:"riderId" validate:"required"`
}

type AssignHubRequest struct {
	HubID string `json:"hubId" validate:"required"`
}
