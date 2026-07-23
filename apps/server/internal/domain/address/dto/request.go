package dto

import "server/internal/models"

type CreateAddressRequest struct {
	Zone        models.ZoneType `json:"zone" validate:"omitempty"`
	District    models.District `json:"district" validate:"required"`
	Division    models.Division `json:"division" validate:"omitempty"`
	FullAddress string          `json:"fullAddress" validate:"required"`
	AreaDetail  string          `json:"areaDetail" validate:"omitempty"`
	PostalCode  string          `json:"postalCode" validate:"omitempty"`
}

type UpdateAddressRequest struct {
	Zone        models.ZoneType `json:"zone" validate:"omitempty"`
	District    models.District `json:"district" validate:"omitempty"`
	Division    models.Division `json:"division" validate:"omitempty"`
	FullAddress string          `json:"fullAddress" validate:"omitempty"`
	AreaDetail  string          `json:"areaDetail" validate:"omitempty"`
	PostalCode  string          `json:"postalCode" validate:"omitempty"`
}
