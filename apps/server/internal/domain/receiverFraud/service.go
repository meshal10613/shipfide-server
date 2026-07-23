package receiverFraud

import (
	"fmt"

	"server/internal/domain/receiverFraud/dto"
	"server/internal/models"
)

type Service interface {
	CheckPhone(phone string) (*dto.ReceiverFraudProfileResponse, error)
	GetByID(id string) (*dto.ReceiverFraudProfileResponse, error)
	UpdateCodBlocked(id string, codBlocked bool) (*dto.ReceiverFraudProfileResponse, error)
	ListProfiles() ([]dto.ReceiverFraudProfileResponse, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CheckPhone(phone string) (*dto.ReceiverFraudProfileResponse, error) {
	profile, err := s.repo.GetByPhone(phone)
	if err != nil {
		return nil, err
	}
	if profile == nil {
		// New default profile
		dummy := &models.ReceiverFraudProfile{
			Phone:       phone,
			FraudScore:  100,
			RiskLevel:   models.RiskLevelNewCustomer,
			CODBlocked:  false,
		}
		res := dto.ToReceiverFraudProfileResponse(dummy)
		return &res, nil
	}

	res := dto.ToReceiverFraudProfileResponse(profile)
	return &res, nil
}

func (s *service) GetByID(id string) (*dto.ReceiverFraudProfileResponse, error) {
	profile, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if profile == nil {
		return nil, fmt.Errorf("receiver fraud profile not found")
	}

	res := dto.ToReceiverFraudProfileResponse(profile)
	return &res, nil
}

func (s *service) UpdateCodBlocked(id string, codBlocked bool) (*dto.ReceiverFraudProfileResponse, error) {
	profile, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if profile == nil {
		return nil, fmt.Errorf("receiver fraud profile not found")
	}

	profile.CODBlocked = codBlocked
	if err := s.repo.Update(profile); err != nil {
		return nil, err
	}

	res := dto.ToReceiverFraudProfileResponse(profile)
	return &res, nil
}

func (s *service) ListProfiles() ([]dto.ReceiverFraudProfileResponse, error) {
	profiles, err := s.repo.List()
	if err != nil {
		return nil, err
	}

	responses := make([]dto.ReceiverFraudProfileResponse, 0, len(profiles))
	for _, p := range profiles {
		responses = append(responses, dto.ToReceiverFraudProfileResponse(&p))
	}
	return responses, nil
}
