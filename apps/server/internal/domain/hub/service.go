package hub

import (
	"errors"

	"github.com/gofiber/fiber/v3"
	"server/internal/domain/hub/dto"
	"server/internal/models"
	querybuilder "server/pkg/queryBuilder"
)

type Service interface {
	CreateHub(req *hubDto.CreateHubRequest) (*hubDto.HubResponse, error)
	GetHubs(c fiber.Ctx, params querybuilder.QueryParams) ([]*hubDto.HubResponse, int64, error)
	GetHub(id string) (*hubDto.HubResponse, error)
	UpdateHub(id string, req *hubDto.UpdateHubRequest) (*hubDto.HubResponse, error)
	DeleteHub(id string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateHub(req *hubDto.CreateHubRequest) (*hubDto.HubResponse, error) {
	hub := &models.Hub{
		Name:       req.Name,
		District:   req.District,
		Division:   req.Division,
		PostalCode: req.PostalCode,
		Address:    req.Address,
	}

	if err := s.repo.Create(hub); err != nil {
		return nil, err
	}

	return hubDto.MapToHubResponse(hub), nil
}

func (s *service) GetHubs(c fiber.Ctx, params querybuilder.QueryParams) ([]*hubDto.HubResponse, int64, error) {
	hubs, total, err := s.repo.FindAll(c, params)
	if err != nil {
		return nil, 0, err
	}
	return hubDto.MapToHubResponseList(hubs), total, nil
}

func (s *service) GetHub(id string) (*hubDto.HubResponse, error) {
	hub, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("hub not found")
	}
	return hubDto.MapToHubResponse(hub), nil
}

func (s *service) UpdateHub(id string, req *hubDto.UpdateHubRequest) (*hubDto.HubResponse, error) {
	hub, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("hub not found")
	}

	if req.Name != nil {
		hub.Name = *req.Name
	}
	if req.District != nil {
		hub.District = *req.District
	}
	if req.Division != nil {
		hub.Division = *req.Division
	}
	if req.PostalCode != nil {
		hub.PostalCode = *req.PostalCode
	}
	if req.Address != nil {
		hub.Address = *req.Address
	}

	if err := s.repo.Update(hub); err != nil {
		return nil, err
	}

	return hubDto.MapToHubResponse(hub), nil
}

func (s *service) DeleteHub(id string) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("hub not found")
	}

	hasStaff, err := s.repo.HasStaff(id)
	if err != nil {
		return err
	}
	if hasStaff {
		return errors.New("cannot delete hub: staff are still assigned. Reassign staff to another hub first")
	}

	return s.repo.Delete(id)
}
