package withdrawal

import (
	"fmt"
	"time"

	"server/internal/domain/withdrawal/dto"
	"server/internal/models"
)

type Service interface {
	CreateWithdrawal(req dto.CreateWithdrawalRequest) (*dto.WithdrawalResponse, error)
	GetWithdrawalByID(id string) (*dto.WithdrawalResponse, error)
	UpdateStatus(id string, req dto.UpdateWithdrawalStatusRequest) (*dto.WithdrawalResponse, error)
	ListWithdrawals(riderID *string, merchantID *string) ([]dto.WithdrawalResponse, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateWithdrawal(req dto.CreateWithdrawalRequest) (*dto.WithdrawalResponse, error) {
	withdrawal := &models.Withdrawal{
		RiderID:               req.RiderID,
		MerchantID:            req.MerchantID,
		Amount:                req.Amount,
		Method:                req.Method,
		AccountNumber:         req.AccountNumber,
		BankName:              req.BankName,
		Status:                models.WithdrawalRequested,
		SendEmailConfirmation: req.SendEmailConfirmation,
	}

	if err := s.repo.Create(withdrawal); err != nil {
		return nil, err
	}

	res := dto.ToWithdrawalResponse(withdrawal)
	return &res, nil
}

func (s *service) GetWithdrawalByID(id string) (*dto.WithdrawalResponse, error) {
	w, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if w == nil {
		return nil, fmt.Errorf("withdrawal request not found")
	}

	res := dto.ToWithdrawalResponse(w)
	return &res, nil
}

func (s *service) UpdateStatus(id string, req dto.UpdateWithdrawalStatusRequest) (*dto.WithdrawalResponse, error) {
	w, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if w == nil {
		return nil, fmt.Errorf("withdrawal request not found")
	}

	w.Status = req.Status
	now := time.Now()

	switch req.Status {
	case models.WithdrawalApproved:
		w.ApprovedAt = &now
	case models.WithdrawalPaid:
		w.PaidAt = &now
	case models.WithdrawalRejected:
		w.RejectionReason = req.RejectionReason
	}

	if err := s.repo.Update(w); err != nil {
		return nil, err
	}

	res := dto.ToWithdrawalResponse(w)
	return &res, nil
}

func (s *service) ListWithdrawals(riderID *string, merchantID *string) ([]dto.WithdrawalResponse, error) {
	items, err := s.repo.List(riderID, merchantID)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.WithdrawalResponse, 0, len(items))
	for _, w := range items {
		responses = append(responses, dto.ToWithdrawalResponse(&w))
	}
	return responses, nil
}
