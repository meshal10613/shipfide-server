package dto

import (
	"time"

	"server/internal/models"
)

type CreateRiderRequest struct {
	UserID                   string                 `json:"userId" validate:"omitempty"`
	Name                     string                 `json:"name" validate:"required"`
	FathersName              string                 `json:"fathersName" validate:"required"`
	MothersName              string                 `json:"mothersName" validate:"required"`
	DateOfBirth              time.Time              `json:"dateOfBirth" validate:"required"`
	Gender                   models.Gender          `json:"gender" validate:"required"`
	BloodGroup               models.BloodGroup      `json:"bloodGroup" validate:"required"`
	District                 models.District        `json:"district" validate:"required"`
	Division                 models.Division        `json:"division" validate:"omitempty"`
	OperatingDistrict        models.District        `json:"operatingDistrict" validate:"required"`
	OperatingDivision        models.Division        `json:"operatingDivision" validate:"omitempty"`
	Address                  string                 `json:"address" validate:"omitempty"`
	EmergencyContactName     string                 `json:"emergencyContactName" validate:"required"`
	EmergencyContactPhone    string                 `json:"emergencyContactPhone" validate:"required"`
	EmergencyContactRelation string                 `json:"emergencyContactRelation" validate:"required"`
	VehicleCategory          models.VehicleCategory `json:"vehicleCategory" validate:"required"`
	VehicleSubType           models.VehicleSubType  `json:"vehicleSubType" validate:"required"`
	Nid                      string                 `json:"nid" validate:"required"`
	DrivingLicense           *string                `json:"drivingLicense" validate:"omitempty"`
	VehicleRC                *string                `json:"vehicleRc" validate:"omitempty"`
	AssignedHubID            *string                `json:"assignedHubId" validate:"omitempty"`
}

type UpdateRiderRequest struct {
	Name                     string                 `json:"name" validate:"omitempty"`
	FathersName              string                 `json:"fathersName" validate:"omitempty"`
	MothersName              string                 `json:"mothersName" validate:"omitempty"`
	District                 models.District        `json:"district" validate:"omitempty"`
	Division                 models.Division        `json:"division" validate:"omitempty"`
	OperatingDistrict        models.District        `json:"operatingDistrict" validate:"omitempty"`
	OperatingDivision        models.Division        `json:"operatingDivision" validate:"omitempty"`
	Address                  string                 `json:"address" validate:"omitempty"`
	EmergencyContactName     string                 `json:"emergencyContactName" validate:"omitempty"`
	EmergencyContactPhone    string                 `json:"emergencyContactPhone" validate:"omitempty"`
	EmergencyContactRelation string                 `json:"emergencyContactRelation" validate:"omitempty"`
	VehicleCategory          models.VehicleCategory `json:"vehicleCategory" validate:"omitempty"`
	VehicleSubType           models.VehicleSubType  `json:"vehicleSubType" validate:"omitempty"`
	Nid                      string                 `json:"nid" validate:"omitempty"`
	DrivingLicense           *string                `json:"drivingLicense" validate:"omitempty"`
	VehicleRC                *string                `json:"vehicleRc" validate:"omitempty"`
	AssignedHubID            *string                `json:"assignedHubId" validate:"omitempty"`
}

type UpdateKycStatusRequest struct {
	KycStatus models.RiderKycStatus `json:"kycStatus" validate:"required"`
}

type UpdateRiderStatusRequest struct {
	IsActive bool `json:"isActive"`
}
