package dto

import (
	"time"

	"server/internal/models"
)

type AddressResponse struct {
	ID          string          `json:"id"`
	Zone        models.ZoneType `json:"zone"`
	Division    models.Division `json:"division"`
	District    models.District `json:"district"`
	FullAddress string          `json:"fullAddress"`
	AreaDetail  string          `json:"areaDetail,omitempty"`
	PostalCode  string          `json:"postalCode,omitempty"`
	CreatedAt   time.Time       `json:"createdAt"`
}

func ToAddressResponse(addr *models.Address) AddressResponse {
	return AddressResponse{
		ID:          addr.ID,
		Zone:        addr.Zone,
		Division:    addr.Division,
		District:    addr.District,
		FullAddress: addr.FullAddress,
		AreaDetail:  addr.AreaDetail,
		PostalCode:  addr.PostalCode,
		CreatedAt:   addr.CreatedAt,
	}
}
