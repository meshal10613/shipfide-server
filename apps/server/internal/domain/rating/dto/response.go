package dto

import (
	"time"

	"server/internal/models"
)

type DeliveryRatingResponse struct {
	ID            string    `json:"id"`
	ShipmentID    string    `json:"shipmentId"`
	RiderID       string    `json:"riderId"`
	ReceiverPhone string    `json:"receiverPhone"`
	IsAnonymous   bool      `json:"isAnonymous"`
	Stars         int       `json:"stars"`
	Comment       *string   `json:"comment,omitempty"`
	SubmittedAt   *time.Time `json:"submittedAt,omitempty"`
	CreatedAt     time.Time `json:"createdAt"`
}

type MerchantRatingResponse struct {
	ID          string                    `json:"id"`
	ShipmentID  string                    `json:"shipmentId"`
	MerchantID  string                    `json:"merchantId"`
	Stars       int                       `json:"stars"`
	Comment     *string                   `json:"comment,omitempty"`
	Tag         *models.MerchantRatingTag `json:"tag,omitempty"`
	SubmittedAt *time.Time                `json:"submittedAt,omitempty"`
	CreatedAt   time.Time                 `json:"createdAt"`
}

func ToDeliveryRatingResponse(d *models.DeliveryRating) DeliveryRatingResponse {
	phone := d.ReceiverPhone
	if d.IsAnonymous {
		phone = "***"
	}
	return DeliveryRatingResponse{
		ID:            d.ID,
		ShipmentID:    d.ShipmentID,
		RiderID:       d.RiderID,
		ReceiverPhone: phone,
		IsAnonymous:   d.IsAnonymous,
		Stars:         d.Stars,
		Comment:       d.Comment,
		SubmittedAt:   d.SubmittedAt,
		CreatedAt:     d.CreatedAt,
	}
}

func ToMerchantRatingResponse(m *models.MerchantDeliveryRating) MerchantRatingResponse {
	return MerchantRatingResponse{
		ID:          m.ID,
		ShipmentID:  m.ShipmentID,
		MerchantID:  m.MerchantID,
		Stars:       m.Stars,
		Comment:     m.Comment,
		Tag:         m.Tag,
		SubmittedAt: m.SubmittedAt,
		CreatedAt:   m.CreatedAt,
	}
}
