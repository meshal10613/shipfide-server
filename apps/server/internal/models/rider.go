package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Rider struct {
	ID     string `gorm:"type:uuid;primaryKey"`
	UserID string `gorm:"type:uuid;not null;uniqueIndex"` // 1-to-1 with User

	Name        string `gorm:"not null"`
	FathersName string `gorm:"not null"`
	MothersName string `gorm:"not null"`
	DateOfBirth time.Time
	Gender      Gender     `gorm:"type:varchar(20);not null"`
	BloodGroup  BloodGroup `gorm:"type:varchar(20);not null"`

	// Home location — Division auto-filled from District.
	Division Division `gorm:"type:varchar(20)"`
	District District `gorm:"type:varchar(60)"`

	// Operational zone — Division auto-filled from District.
	OperatingDivision Division `gorm:"type:varchar(20)"`
	OperatingDistrict District `gorm:"type:varchar(60)"`

	// Full address text (for display).
	Address string `gorm:"type:text"`

	// Emergency contact — all three fields required per spec.
	EmergencyContactName     string `gorm:"not null"`
	EmergencyContactPhone    string `gorm:"type:varchar(20);not null"`
	EmergencyContactRelation string `gorm:"not null"`

	// KYC & vehicle info
	KycStatus       RiderKycStatus  `gorm:"type:varchar(30);default:'PENDING'"`
	VehicleCategory VehicleCategory `gorm:"type:varchar(20)"`
	VehicleSubType  VehicleSubType  `gorm:"type:varchar(20)"`

	// Vehicle document file paths (S3 keys or URLs).
	Nid            string  `gorm:"type:text"` // required
	DrivingLicense *string `gorm:"type:text"`
	VehicleRC      *string `gorm:"type:text"` // required (except bicycle)

	// AssignedHubID is a real FK to Hub (was a bare free-text hub name before,
	// which had no referential integrity and would silently desync on hub renames).
	AssignedHubID *string `gorm:"type:uuid;index"`
	AssignedHub   *Hub    `gorm:"foreignKey:AssignedHubID;references:ID"`

	IsActive bool `gorm:"default:true"` // rider can set false when unavailable

	// Rating cache — updated on every new DeliveryRating submission.
	TotalRatingSum        float64          `gorm:"default:0"`
	TotalRatingCount      int              `gorm:"default:0"`
	AverageRating         float64          `gorm:"default:0"`
	RatingBadge           RiderRatingBadge `gorm:"type:varchar(20);default:'NEW_RIDER'"`
	ConsecutiveLowRatings int              `gorm:"default:0"`

	CreatedAt time.Time
	UpdatedAt time.Time

	// Deleting or updating the parent User cascades to this Rider profile.
	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`

	Shipments       []Shipment       `gorm:"foreignKey:RiderID;references:ID"`
	RatingsReceived []DeliveryRating `gorm:"foreignKey:RiderID;references:ID"`
}

func (r *Rider) validateDivisions() error {
	if r.District != "" {
		exp, ok := DivisionOf(r.District)
		if !ok {
			return fmt.Errorf("unknown district: %q", r.District)
		}
		if r.Division != "" && r.Division != exp {
			return fmt.Errorf("district %q belongs to division %q, not %q", r.District, exp, r.Division)
		}
		r.Division = exp
	}
	if r.OperatingDistrict != "" {
		exp, ok := DivisionOf(r.OperatingDistrict)
		if !ok {
			return fmt.Errorf("unknown operating district: %q", r.OperatingDistrict)
		}
		if r.OperatingDivision != "" && r.OperatingDivision != exp {
			return fmt.Errorf("operating district %q belongs to division %q, not %q", r.OperatingDistrict, exp, r.OperatingDivision)
		}
		r.OperatingDivision = exp
	}
	return nil
}

func (r *Rider) BeforeCreate(tx *gorm.DB) error {
	if r.ID == "" {
		r.ID = uuid.NewString()
	}
	return r.validateDivisions()
}

func (r *Rider) BeforeSave(_ *gorm.DB) error {
	return r.validateDivisions()
}
