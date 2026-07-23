package deliveryConfirmation

import (
	"fmt"
	"time"

	"server/internal/domain/deliveryConfirmation/dto"
	"server/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	VerifyOtp(req dto.VerifyOtpRequest) (*dto.DeliveryConfirmationResponse, error)
	RegenerateOtp(adminID string, req dto.RegenerateOtpRequest) (*dto.DeliveryConfirmationResponse, error)
	ConfirmCod(req dto.ConfirmCodRequest) (*dto.CodConfirmationResponse, error)
	DepositCod(req dto.DepositCodRequest) ([]dto.CodConfirmationResponse, error)
	ApproveDeposit(adminID string, codID string) (*dto.CodConfirmationResponse, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) VerifyOtp(req dto.VerifyOtpRequest) (*dto.DeliveryConfirmationResponse, error) {
	dc, err := s.repo.GetOtpByShipmentID(req.ShipmentID)
	if err != nil {
		return nil, err
	}
	if dc == nil {
		return nil, fmt.Errorf("delivery confirmation OTP not found for shipment")
	}

	if dc.IsLocked {
		return nil, fmt.Errorf("OTP confirmation is locked due to too many failed attempts")
	}

	if dc.IsUsed {
		return nil, fmt.Errorf("OTP has already been used")
	}

	if time.Now().After(dc.OtpExpiresAt) {
		return nil, fmt.Errorf("OTP has expired")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dc.OtpHash), []byte(req.Otp)); err != nil {
		dc.WrongAttempts++
		if dc.WrongAttempts >= 3 {
			dc.IsLocked = true
		}
		_ = s.repo.UpdateOtp(dc)
		return nil, fmt.Errorf("invalid OTP")
	}

	now := time.Now()
	dc.IsUsed = true
	dc.UsedAt = &now

	if err := s.repo.UpdateOtp(dc); err != nil {
		return nil, err
	}

	res := dto.ToDeliveryConfirmationResponse(dc)
	return &res, nil
}

func (s *service) RegenerateOtp(adminID string, req dto.RegenerateOtpRequest) (*dto.DeliveryConfirmationResponse, error) {
	dc, err := s.repo.GetOtpByShipmentID(req.ShipmentID)
	if err != nil {
		return nil, err
	}
	if dc == nil {
		return nil, fmt.Errorf("delivery confirmation record not found")
	}

	// Generate new OTP
	newOtp := "123456" // Default fixed OTP for admin override in dev/staging
	hash, err := bcrypt.GenerateFromPassword([]byte(newOtp), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	dc.OtpHash = string(hash)
	dc.OtpExpiresAt = now.Add(48 * time.Hour)
	dc.WrongAttempts = 0
	dc.IsLocked = false
	dc.RegeneratedAt = &now
	dc.RegeneratedByAdmin = &adminID

	if err := s.repo.UpdateOtp(dc); err != nil {
		return nil, err
	}

	res := dto.ToDeliveryConfirmationResponse(dc)
	return &res, nil
}

func (s *service) ConfirmCod(req dto.ConfirmCodRequest) (*dto.CodConfirmationResponse, error) {
	var shipment models.Shipment
	if err := s.repo.DB().First(&shipment, "id = ?", req.ShipmentID).Error; err != nil {
		return nil, fmt.Errorf("shipment not found")
	}

	if !req.CollectedAmount.Equal(shipment.CodAmount) {
		return nil, fmt.Errorf("collected amount %s does not match expected COD amount %s", req.CollectedAmount.String(), shipment.CodAmount.String())
	}

	cod, err := s.repo.GetCodByShipmentID(req.ShipmentID)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	if cod == nil {
		cod = &models.CodDeliveryConfirmation{
			ShipmentID:      req.ShipmentID,
			ExpectedAmount:  shipment.CodAmount,
			CollectedAmount: &req.CollectedAmount,
			ConfirmedAt:     &now,
		}
		if err := s.repo.CreateCod(cod); err != nil {
			return nil, err
		}
	} else {
		cod.CollectedAmount = &req.CollectedAmount
		cod.ConfirmedAt = &now
		if err := s.repo.UpdateCod(cod); err != nil {
			return nil, err
		}
	}

	res := dto.ToCodConfirmationResponse(cod)
	return &res, nil
}

func (s *service) DepositCod(req dto.DepositCodRequest) ([]dto.CodConfirmationResponse, error) {
	now := time.Now()
	responses := make([]dto.CodConfirmationResponse, 0, len(req.CodConfirmationIDs))

	for _, id := range req.CodConfirmationIDs {
		cod, err := s.repo.GetCodByID(id)
		if err != nil || cod == nil {
			continue
		}
		cod.DepositedAt = &now
		_ = s.repo.UpdateCod(cod)
		responses = append(responses, dto.ToCodConfirmationResponse(cod))
	}

	return responses, nil
}

func (s *service) ApproveDeposit(adminID string, codID string) (*dto.CodConfirmationResponse, error) {
	cod, err := s.repo.GetCodByID(codID)
	if err != nil {
		return nil, err
	}
	if cod == nil {
		return nil, fmt.Errorf("COD confirmation record not found")
	}

	now := time.Now()
	cod.DepositConfirmedByAdminID = &adminID
	cod.DepositConfirmedAt = &now

	if err := s.repo.UpdateCod(cod); err != nil {
		return nil, err
	}

	res := dto.ToCodConfirmationResponse(cod)
	return &res, nil
}
