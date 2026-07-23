package dto

import (
	"time"

	"server/internal/models"
)

type RiderResponse struct {
	ID                       string                 `json:"id"`
	UserID                   string                 `json:"userId"`
	Name                     string                 `json:"name"`
	FathersName              string                 `json:"fathersName"`
	MothersName              string                 `json:"mothersName"`
	DateOfBirth              time.Time              `json:"dateOfBirth"`
	Gender                   models.Gender          `json:"gender"`
	BloodGroup               models.BloodGroup      `json:"bloodGroup"`
	Division                 models.Division        `json:"division"`
	District                 models.District        `json:"district"`
	OperatingDivision        models.Division        `json:"operatingDivision"`
	OperatingDistrict        models.District        `json:"operatingDistrict"`
	Address                  string                 `json:"address,omitempty"`
	EmergencyContactName     string                 `json:"emergencyContactName"`
	EmergencyContactPhone    string                 `json:"emergencyContactPhone"`
	EmergencyContactRelation string                 `json:"emergencyContactRelation"`
	KycStatus                models.RiderKycStatus  `json:"kycStatus"`
	VehicleCategory          models.VehicleCategory `json:"vehicleCategory"`
	VehicleSubType           models.VehicleSubType  `json:"vehicleSubType"`
	Nid                      string                 `json:"nid,omitempty"`
	DrivingLicense           *string                `json:"drivingLicense,omitempty"`
	VehicleRC                *string                `json:"vehicleRc,omitempty"`
	AssignedHubID            *string                `json:"assignedHubId,omitempty"`
	AssignedHubName          string                 `json:"assignedHubName,omitempty"`
	IsActive                 bool                   `json:"isActive"`
	TotalRatingSum           float64                `json:"totalRatingSum"`
	TotalRatingCount         int                    `json:"totalRatingCount"`
	AverageRating            float64                `json:"averageRating"`
	RatingBadge              models.RiderRatingBadge `json:"ratingBadge"`
	ConsecutiveLowRatings    int                    `json:"consecutiveLowRatings"`
	CreatedAt                time.Time              `json:"createdAt"`
	UpdatedAt                time.Time              `json:"updatedAt"`
}

func ToRiderResponse(r *models.Rider) RiderResponse {
	hubName := ""
	if r.AssignedHub != nil {
		hubName = r.AssignedHub.Name
	}
	return RiderResponse{
		ID:                       r.ID,
		UserID:                   r.UserID,
		Name:                     r.Name,
		FathersName:              r.FathersName,
		MothersName:              r.MothersName,
		DateOfBirth:              r.DateOfBirth,
		Gender:                   r.Gender,
		BloodGroup:               r.BloodGroup,
		Division:                 r.Division,
		District:                 r.District,
		OperatingDivision:        r.OperatingDivision,
		OperatingDistrict:        r.OperatingDistrict,
		Address:                  r.Address,
		EmergencyContactName:     r.EmergencyContactName,
		EmergencyContactPhone:    r.EmergencyContactPhone,
		EmergencyContactRelation: r.EmergencyContactRelation,
		KycStatus:                r.KycStatus,
		VehicleCategory:          r.VehicleCategory,
		VehicleSubType:           r.VehicleSubType,
		Nid:                      r.Nid,
		DrivingLicense:           r.DrivingLicense,
		VehicleRC:                r.VehicleRC,
		AssignedHubID:            r.AssignedHubID,
		AssignedHubName:          hubName,
		IsActive:                 r.IsActive,
		TotalRatingSum:           r.TotalRatingSum,
		TotalRatingCount:         r.TotalRatingCount,
		AverageRating:            r.AverageRating,
		RatingBadge:              r.RatingBadge,
		ConsecutiveLowRatings:    r.ConsecutiveLowRatings,
		CreatedAt:                r.CreatedAt,
		UpdatedAt:                r.UpdatedAt,
	}
}
