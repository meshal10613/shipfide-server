package dto

import (
	"time"

	"server/internal/models"
)

type GuestSenderResponse struct {
	ID             string          `json:"id"`
	Name           string          `json:"name"`
	Phone          string          `json:"phone"`
	Email          *string         `json:"email,omitempty"`
	Division       models.Division `json:"division,omitempty"`
	District       models.District `json:"district,omitempty"`
	Address        *string         `json:"address,omitempty"`
	NidNumber      *string         `json:"nidNumber,omitempty"`
	IsPhoneFlagged bool            `json:"isPhoneFlagged"`
	FlagReason     string          `json:"flagReason,omitempty"`
	AdminID        string          `json:"adminId"`
	CreatedAt      time.Time       `json:"createdAt"`
}

func ToGuestSenderResponse(g *models.GuestSender) GuestSenderResponse {
	return GuestSenderResponse{
		ID:             g.ID,
		Name:           g.Name,
		Phone:          g.Phone,
		Email:          g.Email,
		Division:       g.Division,
		District:       g.District,
		Address:        g.Address,
		NidNumber:      g.NidNumber,
		IsPhoneFlagged: g.IsPhoneFlagged,
		FlagReason:     g.FlagReason,
		AdminID:        g.AdminID,
		CreatedAt:      g.CreatedAt,
	}
}
