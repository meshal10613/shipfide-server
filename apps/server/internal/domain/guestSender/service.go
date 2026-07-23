package guestSender

import (
	"fmt"

	"server/internal/domain/guestSender/dto"
	"server/internal/models"
)

type Service interface {
	CreateGuestSender(adminID string, req dto.CreateGuestSenderRequest) (*dto.GuestSenderResponse, error)
	GetGuestSenderByID(id string) (*dto.GuestSenderResponse, error)
	FlagGuestSender(id string, req dto.FlagGuestSenderRequest) (*dto.GuestSenderResponse, error)
	ListGuestSenders() ([]dto.GuestSenderResponse, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateGuestSender(adminID string, req dto.CreateGuestSenderRequest) (*dto.GuestSenderResponse, error) {
	existing, err := s.repo.GetByPhone(req.Phone)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		res := dto.ToGuestSenderResponse(existing)
		return &res, nil
	}

	g := &models.GuestSender{
		Name:      req.Name,
		Phone:     req.Phone,
		Email:     req.Email,
		District:  req.District,
		Division:  req.Division,
		Address:   req.Address,
		NidNumber: req.NidNumber,
		AdminID:   adminID,
	}

	if err := s.repo.Create(g); err != nil {
		return nil, err
	}

	res := dto.ToGuestSenderResponse(g)
	return &res, nil
}

func (s *service) GetGuestSenderByID(id string) (*dto.GuestSenderResponse, error) {
	g, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if g == nil {
		return nil, fmt.Errorf("guest sender not found")
	}

	res := dto.ToGuestSenderResponse(g)
	return &res, nil
}

func (s *service) FlagGuestSender(id string, req dto.FlagGuestSenderRequest) (*dto.GuestSenderResponse, error) {
	g, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if g == nil {
		return nil, fmt.Errorf("guest sender not found")
	}

	g.IsPhoneFlagged = req.IsPhoneFlagged
	g.FlagReason = req.FlagReason

	if err := s.repo.Update(g); err != nil {
		return nil, err
	}

	res := dto.ToGuestSenderResponse(g)
	return &res, nil
}

func (s *service) ListGuestSenders() ([]dto.GuestSenderResponse, error) {
	guests, err := s.repo.List()
	if err != nil {
		return nil, err
	}

	responses := make([]dto.GuestSenderResponse, 0, len(guests))
	for _, g := range guests {
		responses = append(responses, dto.ToGuestSenderResponse(&g))
	}
	return responses, nil
}
