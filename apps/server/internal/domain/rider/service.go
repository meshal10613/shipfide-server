package rider

import (
	"context"
	"fmt"

	"server/internal/config"
	"server/internal/domain/rider/dto"
	"server/internal/models"
	cldService "server/pkg/cloudinary"
)

type Service interface {
	CreateRider(userID string, req dto.CreateRiderRequest) (*dto.RiderResponse, error)
	GetRiderByID(id string) (*dto.RiderResponse, error)
	GetRiderByUserID(userID string) (*dto.RiderResponse, error)
	UpdateRider(id string, req dto.UpdateRiderRequest) (*dto.RiderResponse, error)
	UpdateKycStatus(id string, req dto.UpdateKycStatusRequest) (*dto.RiderResponse, error)
	UpdateStatus(id string, req dto.UpdateRiderStatusRequest) (*dto.RiderResponse, error)
	ListRiders() ([]dto.RiderResponse, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateRider(userID string, req dto.CreateRiderRequest) (*dto.RiderResponse, error) {
	existing, err := s.repo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, fmt.Errorf("rider profile already exists for this user")
	}

	rider := &models.Rider{
		UserID:                   userID,
		Name:                     req.Name,
		FathersName:              req.FathersName,
		MothersName:              req.MothersName,
		DateOfBirth:              req.DateOfBirth,
		Gender:                   req.Gender,
		BloodGroup:               req.BloodGroup,
		District:                 req.District,
		Division:                 req.Division,
		OperatingDistrict:        req.OperatingDistrict,
		OperatingDivision:        req.OperatingDivision,
		Address:                  req.Address,
		EmergencyContactName:     req.EmergencyContactName,
		EmergencyContactPhone:    req.EmergencyContactPhone,
		EmergencyContactRelation: req.EmergencyContactRelation,
		VehicleCategory:          req.VehicleCategory,
		VehicleSubType:           req.VehicleSubType,
		Nid:                      req.Nid,
		DrivingLicense:           req.DrivingLicense,
		VehicleRC:                req.VehicleRC,
		AssignedHubID:            req.AssignedHubID,
		KycStatus:                models.RiderKycPendingReview,
		IsActive:                 true,
		RatingBadge:              models.RiderBadgeNewRider,
	}

	if err := s.repo.Create(rider); err != nil {
		return nil, err
	}

	// Reload for relations
	updatedRider, _ := s.repo.GetByID(rider.ID)
	if updatedRider != nil {
		rider = updatedRider
	}

	res := dto.ToRiderResponse(rider)
	return &res, nil
}

func (s *service) GetRiderByID(id string) (*dto.RiderResponse, error) {
	rider, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if rider == nil {
		return nil, fmt.Errorf("rider not found")
	}

	res := dto.ToRiderResponse(rider)
	return &res, nil
}

func (s *service) GetRiderByUserID(userID string) (*dto.RiderResponse, error) {
	rider, err := s.repo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}
	if rider == nil {
		return nil, fmt.Errorf("rider profile not found for user")
	}

	res := dto.ToRiderResponse(rider)
	return &res, nil
}

func (s *service) UpdateRider(id string, req dto.UpdateRiderRequest) (*dto.RiderResponse, error) {
	rider, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if rider == nil {
		return nil, fmt.Errorf("rider not found")
	}

	oldNid := rider.Nid
	var oldDrivingLicense, oldVehicleRC string
	if rider.DrivingLicense != nil {
		oldDrivingLicense = *rider.DrivingLicense
	}
	if rider.VehicleRC != nil {
		oldVehicleRC = *rider.VehicleRC
	}

	if req.Name != "" {
		rider.Name = req.Name
	}
	if req.FathersName != "" {
		rider.FathersName = req.FathersName
	}
	if req.MothersName != "" {
		rider.MothersName = req.MothersName
	}
	if req.District != "" {
		rider.District = req.District
	}
	if req.Division != "" {
		rider.Division = req.Division
	}
	if req.OperatingDistrict != "" {
		rider.OperatingDistrict = req.OperatingDistrict
	}
	if req.OperatingDivision != "" {
		rider.OperatingDivision = req.OperatingDivision
	}
	if req.Address != "" {
		rider.Address = req.Address
	}
	if req.EmergencyContactName != "" {
		rider.EmergencyContactName = req.EmergencyContactName
	}
	if req.EmergencyContactPhone != "" {
		rider.EmergencyContactPhone = req.EmergencyContactPhone
	}
	if req.EmergencyContactRelation != "" {
		rider.EmergencyContactRelation = req.EmergencyContactRelation
	}
	if req.VehicleCategory != "" {
		rider.VehicleCategory = req.VehicleCategory
	}
	if req.VehicleSubType != "" {
		rider.VehicleSubType = req.VehicleSubType
	}
	if req.Nid != "" {
		rider.Nid = req.Nid
	}
	if req.DrivingLicense != nil {
		rider.DrivingLicense = req.DrivingLicense
	}
	if req.VehicleRC != nil {
		rider.VehicleRC = req.VehicleRC
	}
	if req.AssignedHubID != nil {
		rider.AssignedHubID = req.AssignedHubID
	}

	if err := s.repo.Update(rider); err != nil {
		return nil, err
	}

	// Clean up old Cloudinary images when replaced
	cld, _ := cldService.NewCloudinaryService(config.AppConfig)
	if cld != nil {
		if req.Nid != "" && oldNid != "" && oldNid != req.Nid {
			_ = cld.DeleteImage(context.Background(), oldNid)
		}
		if req.DrivingLicense != nil && oldDrivingLicense != "" && oldDrivingLicense != *req.DrivingLicense {
			_ = cld.DeleteImage(context.Background(), oldDrivingLicense)
		}
		if req.VehicleRC != nil && oldVehicleRC != "" && oldVehicleRC != *req.VehicleRC {
			_ = cld.DeleteImage(context.Background(), oldVehicleRC)
		}
	}

	updatedRider, _ := s.repo.GetByID(rider.ID)
	if updatedRider != nil {
		rider = updatedRider
	}

	res := dto.ToRiderResponse(rider)
	return &res, nil
}

func (s *service) UpdateKycStatus(id string, req dto.UpdateKycStatusRequest) (*dto.RiderResponse, error) {
	rider, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if rider == nil {
		return nil, fmt.Errorf("rider not found")
	}

	rider.KycStatus = req.KycStatus
	if err := s.repo.Update(rider); err != nil {
		return nil, err
	}

	res := dto.ToRiderResponse(rider)
	return &res, nil
}

func (s *service) UpdateStatus(id string, req dto.UpdateRiderStatusRequest) (*dto.RiderResponse, error) {
	rider, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if rider == nil {
		return nil, fmt.Errorf("rider not found")
	}

	rider.IsActive = req.IsActive
	if err := s.repo.Update(rider); err != nil {
		return nil, err
	}

	res := dto.ToRiderResponse(rider)
	return &res, nil
}

func (s *service) ListRiders() ([]dto.RiderResponse, error) {
	riders, err := s.repo.List()
	if err != nil {
		return nil, err
	}

	responses := make([]dto.RiderResponse, 0, len(riders))
	for _, r := range riders {
		responses = append(responses, dto.ToRiderResponse(&r))
	}
	return responses, nil
}
