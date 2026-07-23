package walkInPayment

import (
	"fmt"
	"time"

	"server/internal/domain/walkInPayment/dto"
	"server/internal/models"
)

type Service interface {
	CreatePayment(adminID string, req dto.CreateWalkInPaymentRequest) (*dto.WalkInPaymentResponse, error)
	ListPayments() ([]dto.WalkInPaymentResponse, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreatePayment(adminID string, req dto.CreateWalkInPaymentRequest) (*dto.WalkInPaymentResponse, error) {
	payment := &models.WalkInPayment{
		ShipmentID:         req.ShipmentID,
		Amount:             req.Amount,
		Method:             req.Method,
		Status:             models.WalkInCollected,
		CollectedByAdminID: adminID,
		CollectedAt:        time.Now(),
	}

	if err := s.repo.Create(payment); err != nil {
		return nil, fmt.Errorf("failed to record walk-in payment: %w", err)
	}

	res := dto.ToWalkInPaymentResponse(payment)
	return &res, nil
}

func (s *service) ListPayments() ([]dto.WalkInPaymentResponse, error) {
	payments, err := s.repo.List()
	if err != nil {
		return nil, err
	}

	responses := make([]dto.WalkInPaymentResponse, 0, len(payments))
	for _, p := range payments {
		responses = append(responses, dto.ToWalkInPaymentResponse(&p))
	}
	return responses, nil
}
