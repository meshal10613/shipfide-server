package address

import (
	"fmt"

	"server/internal/domain/address/dto"
	"server/internal/models"
)

type Service interface {
	CreateAddress(req dto.CreateAddressRequest) (*dto.AddressResponse, error)
	GetAddressByID(id string) (*dto.AddressResponse, error)
	ListAddresses() ([]dto.AddressResponse, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateAddress(req dto.CreateAddressRequest) (*dto.AddressResponse, error) {
	addr := &models.Address{
		Zone:        req.Zone,
		District:    req.District,
		Division:    req.Division,
		FullAddress: req.FullAddress,
		AreaDetail:  req.AreaDetail,
		PostalCode:  req.PostalCode,
	}

	if err := s.repo.Create(addr); err != nil {
		return nil, err
	}

	res := dto.ToAddressResponse(addr)
	return &res, nil
}

func (s *service) GetAddressByID(id string) (*dto.AddressResponse, error) {
	addr, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if addr == nil {
		return nil, fmt.Errorf("address not found")
	}

	res := dto.ToAddressResponse(addr)
	return &res, nil
}

func (s *service) ListAddresses() ([]dto.AddressResponse, error) {
	addresses, err := s.repo.List()
	if err != nil {
		return nil, err
	}

	responses := make([]dto.AddressResponse, 0, len(addresses))
	for _, addr := range addresses {
		responses = append(responses, dto.ToAddressResponse(&addr))
	}
	return responses, nil
}
