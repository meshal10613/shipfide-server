package dto

import (
	"time"

	"server/internal/models"

	"github.com/shopspring/decimal"
)

type ShipmentResponse struct {
	ID                 string                    `json:"id"`
	TrackingCode       string                    `json:"trackingCode"`
	SenderType         models.SenderType         `json:"senderType"`
	MerchantID         *string                   `json:"merchantId,omitempty"`
	GuestSenderID      *string                   `json:"guestSenderId,omitempty"`
	CreatedByAdminID   *string                   `json:"createdByAdminId,omitempty"`
	RiderID            *string                   `json:"riderId,omitempty"`
	HubID              *string                   `json:"hubId,omitempty"`
	Status             models.ShipmentStatus     `json:"status"`
	ParcelType         models.ParcelType         `json:"parcelType"`
	Weight             int                       `json:"weight"`
	Dimensions         *string                   `json:"dimensions,omitempty"`
	ProductName        string                    `json:"productName"`
	ProductValue       *decimal.Decimal          `json:"productValue,omitempty"`
	IsFragile          bool                      `json:"isFragile"`
	Notes              *string                   `json:"notes,omitempty"`
	SenderName         string                    `json:"senderName"`
	SenderPhone        string                    `json:"senderPhone"`
	SenderAddress      *string                   `json:"senderAddress,omitempty"`
	ReceiverName       string                    `json:"receiverName"`
	ReceiverPhone      string                    `json:"receiverPhone"`
	ReceiverEmail      *string                   `json:"receiverEmail,omitempty"`
	ReceiverAddress    string                    `json:"receiverAddress"`
	ReceiverAddressID  *string                   `json:"receiverAddressId,omitempty"`
	DeliveryCharge     decimal.Decimal           `json:"deliveryCharge"`
	CodAmount          decimal.Decimal           `json:"codAmount"`
	TotalCharge        decimal.Decimal           `json:"totalCharge"`
	ConfirmationMethod models.ConfirmationMethod `json:"confirmationMethod"`
	SplitZone          *models.SplitZone         `json:"splitZone,omitempty"`
	RiderSharePct      *decimal.Decimal          `json:"riderSharePct,omitempty"`
	RiderShare         *decimal.Decimal          `json:"riderShare,omitempty"`
	SystemShare        *decimal.Decimal          `json:"systemShare,omitempty"`
	MerchantNet        *decimal.Decimal          `json:"merchantNet,omitempty"`
	DeliveredAt        *time.Time                `json:"deliveredAt,omitempty"`
	CreatedAt          time.Time                 `json:"createdAt"`
	UpdatedAt          time.Time                 `json:"updatedAt"`
}

func ToShipmentResponse(s *models.Shipment) ShipmentResponse {
	return ShipmentResponse{
		ID:                 s.ID,
		TrackingCode:       s.TrackingCode,
		SenderType:         s.SenderType,
		MerchantID:         s.MerchantID,
		GuestSenderID:      s.GuestSenderID,
		CreatedByAdminID:   s.CreatedByAdminID,
		RiderID:            s.RiderID,
		HubID:              s.HubID,
		Status:             s.Status,
		ParcelType:         s.ParcelType,
		Weight:             s.Weight,
		Dimensions:         s.Dimensions,
		ProductName:        s.ProductName,
		ProductValue:       s.ProductValue,
		IsFragile:          s.IsFragile,
		Notes:              s.Notes,
		SenderName:         s.SenderName,
		SenderPhone:        s.SenderPhone,
		SenderAddress:      s.SenderAddress,
		ReceiverName:       s.ReceiverName,
		ReceiverPhone:      s.ReceiverPhone,
		ReceiverEmail:      s.ReceiverEmail,
		ReceiverAddress:    s.ReceiverAddress,
		ReceiverAddressID:  s.ReceiverAddressID,
		DeliveryCharge:     s.DeliveryCharge,
		CodAmount:          s.CodAmount,
		TotalCharge:        s.TotalCharge,
		ConfirmationMethod: s.ConfirmationMethod,
		SplitZone:          s.SplitZone,
		RiderSharePct:      s.RiderSharePct,
		RiderShare:         s.RiderShare,
		SystemShare:        s.SystemShare,
		MerchantNet:        s.MerchantNet,
		DeliveredAt:        s.DeliveredAt,
		CreatedAt:          s.CreatedAt,
		UpdatedAt:          s.UpdatedAt,
	}
}
