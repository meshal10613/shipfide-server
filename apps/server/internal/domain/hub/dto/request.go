package hubDto

import (
	"server/internal/models"
)

type CreateHubRequest struct {
	Name       string          `json:"name" validate:"required,min=3,max=120"`
	District   models.District `json:"district" validate:"required"`
	Division   models.Division `json:"division" validate:"omitempty"`
	PostalCode string          `json:"postalCode" validate:"required,min=3,max=10"`
	Address    string          `json:"address" validate:"required,min=5"`
}

type UpdateHubRequest struct {
	Name       *string          `json:"name" validate:"omitempty,min=3,max=120"`
	District   *models.District `json:"district" validate:"omitempty"`
	Division   *models.Division `json:"division" validate:"omitempty"`
	PostalCode *string          `json:"postalCode" validate:"omitempty,min=3,max=10"`
	Address    *string          `json:"address" validate:"omitempty,min=5"`
}
