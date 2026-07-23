package merchant

import (
	"fmt"

	"server/internal/domain/merchant/dto"
	"server/internal/models"

	"github.com/shopspring/decimal"
)

type Service interface {
	CreateMerchant(userID string, req dto.CreateMerchantRequest) (*dto.MerchantResponse, error)
	GetMerchantByID(id string) (*dto.MerchantResponse, error)
	GetMerchantByUserID(userID string) (*dto.MerchantResponse, error)
	UpdateMerchant(id string, req dto.UpdateMerchantRequest) (*dto.MerchantResponse, error)
	UpdateKyc(id string, req dto.UpdateKycRequest) (*dto.MerchantResponse, error)
	ListMerchants() ([]dto.MerchantResponse, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateMerchant(userID string, req dto.CreateMerchantRequest) (*dto.MerchantResponse, error) {
	existing, err := s.repo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, fmt.Errorf("merchant profile already exists for this user")
	}

	maxCod := decimal.NewFromInt(50000)

	m := &models.Merchant{
		UserID:         userID,
		Name:           req.Name,
		FathersName:    req.FathersName,
		MothersName:    req.MothersName,
		DateOfBirth:    req.DateOfBirth,
		Gender:         req.Gender,
		BloodGroup:     req.BloodGroup,
		District:       req.District,
		Division:       req.Division,
		PickupDistrict: req.PickupDistrict,
		PickupDivision: req.PickupDivision,
		Address:        req.Address,
		Nid:            req.Nid,
		CodEnabled:     true,
		MaxCodAmount:   maxCod,
	}

	if err := s.repo.Create(m); err != nil {
		return nil, err
	}

	res := dto.ToMerchantResponse(m)
	return &res, nil
}

func (s *service) GetMerchantByID(id string) (*dto.MerchantResponse, error) {
	m, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if m == nil {
		return nil, fmt.Errorf("merchant not found")
	}

	res := dto.ToMerchantResponse(m)
	return &res, nil
}

func (s *service) GetMerchantByUserID(userID string) (*dto.MerchantResponse, error) {
	m, err := s.repo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}
	if m == nil {
		return nil, fmt.Errorf("merchant profile not found for user")
	}

	res := dto.ToMerchantResponse(m)
	return &res, nil
}

func (s *service) UpdateMerchant(id string, req dto.UpdateMerchantRequest) (*dto.MerchantResponse, error) {
	m, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if m == nil {
		return nil, fmt.Errorf("merchant not found")
	}

	if req.Name != "" {
		m.Name = req.Name
	}
	if req.FathersName != "" {
		m.FathersName = req.FathersName
	}
	if req.MothersName != "" {
		m.MothersName = req.MothersName
	}
	if req.District != "" {
		m.District = req.District
	}
	if req.Division != "" {
		m.Division = req.Division
	}
	if req.PickupDistrict != "" {
		m.PickupDistrict = req.PickupDistrict
	}
	if req.PickupDivision != "" {
		m.PickupDivision = req.PickupDivision
	}
	if req.Address != "" {
		m.Address = req.Address
	}
	if req.Nid != "" {
		m.Nid = req.Nid
	}

	if err := s.repo.Update(m); err != nil {
		return nil, err
	}

	res := dto.ToMerchantResponse(m)
	return &res, nil
}

func (s *service) UpdateKyc(id string, req dto.UpdateKycRequest) (*dto.MerchantResponse, error) {
	m, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if m == nil {
		return nil, fmt.Errorf("merchant not found")
	}

	m.IsKycVerified = req.IsKycVerified
	if req.CodEnabled != nil {
		m.CodEnabled = *req.CodEnabled
	}
	if req.MaxCodAmount != nil {
		m.MaxCodAmount = *req.MaxCodAmount
	}

	if err := s.repo.Update(m); err != nil {
		return nil, err
	}

	res := dto.ToMerchantResponse(m)
	return &res, nil
}

func (s *service) ListMerchants() ([]dto.MerchantResponse, error) {
	merchants, err := s.repo.List()
	if err != nil {
		return nil, err
	}

	responses := make([]dto.MerchantResponse, 0, len(merchants))
	for _, m := range merchants {
		responses = append(responses, dto.ToMerchantResponse(&m))
	}
	return responses, nil
}
