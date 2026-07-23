package hubDto

import (
	"time"
	"server/internal/models"
)

type HubResponse struct {
	ID         string          `json:"id"`
	Name       string          `json:"name"`
	Division   models.Division `json:"division"`
	District   models.District `json:"district"`
	PostalCode string          `json:"postalCode"`
	Address    string          `json:"address"`
	CreatedAt  time.Time       `json:"createdAt"`
	UpdatedAt  time.Time       `json:"updatedAt"`
}

func MapToHubResponse(h *models.Hub) *HubResponse {
	return &HubResponse{
		ID:         h.ID,
		Name:       h.Name,
		Division:   h.Division,
		District:   h.District,
		PostalCode: h.PostalCode,
		Address:    h.Address,
		CreatedAt:  h.CreatedAt,
		UpdatedAt:  h.UpdatedAt,
	}
}

func MapToHubResponseList(hubs []models.Hub) []*HubResponse {
	res := make([]*HubResponse, len(hubs))
	for i, h := range hubs {
		res[i] = MapToHubResponse(&h)
	}
	return res
}
