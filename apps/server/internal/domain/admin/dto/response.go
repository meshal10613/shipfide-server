package adminDto

import (
	"time"
	"server/internal/models"
)

type AdminResponse struct {
	ID                string          `json:"id"`
	UserID            string          `json:"userId"`
	Email             string          `json:"email"`
	Name              string          `json:"name"`
	Phone             string          `json:"phone"`
	Division          models.Division `json:"division"`
	District          models.District `json:"district"`
	HubID             *string         `json:"hubId,omitempty"`
	CanApproveCashout bool            `json:"canApproveCashout"`
	CanManageRiders   bool            `json:"canManageRiders"`
	CanViewFraudFlags bool            `json:"canViewFraudFlags"`
	CreatedAt         time.Time       `json:"createdAt"`
	UpdatedAt         time.Time       `json:"updatedAt"`
}

func MapToAdminResponse(a *models.Admin) *AdminResponse {
	email := ""
	if a.User != nil {
		email = a.User.Email
	}
	return &AdminResponse{
		ID:                a.ID,
		UserID:            a.UserID,
		Email:             email,
		Name:              a.Name,
		Phone:             a.Phone,
		Division:          a.Division,
		District:          a.District,
		HubID:             a.HubID,
		CanApproveCashout: a.CanApproveCashout,
		CanManageRiders:   a.CanManageRiders,
		CanViewFraudFlags: a.CanViewFraudFlags,
		CreatedAt:         a.CreatedAt,
		UpdatedAt:         a.UpdatedAt,
	}
}

func MapToAdminResponseList(admins []models.Admin) []*AdminResponse {
	res := make([]*AdminResponse, len(admins))
	for i, a := range admins {
		res[i] = MapToAdminResponse(&a)
	}
	return res
}
