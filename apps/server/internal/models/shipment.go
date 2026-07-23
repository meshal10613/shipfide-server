package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Shipment struct {
	ID               string     `gorm:"type:uuid;primaryKey"`
	TrackingCode     string     `gorm:"not null;uniqueIndex"`
	SenderType       SenderType `gorm:"type:varchar(20);not null;default:'GUEST'"`
	MerchantID       *string    `gorm:"type:uuid;index"`
	GuestSenderID    *string    `gorm:"type:uuid;index"`
	CreatedByAdminID *string    `gorm:"type:uuid;index"`
	RiderID          *string    `gorm:"type:uuid;index"`
	HubID            *string    `gorm:"type:uuid;index"`

	// Status carries its own index (in addition to the composite ones below)
	// because "list all shipments in state X across every hub" is a common
	// admin/ops query on its own.
	Status       ShipmentStatus   `gorm:"type:varchar(40);not null;default:'PENDING';index"`
	ParcelType   ParcelType       `gorm:"type:varchar(20);not null;default:'PACKAGE'"`
	Weight       int              `gorm:"not null"` // grams
	Dimensions   *string          `gorm:"type:text"`
	ProductName  string           `gorm:"not null"`
	ProductValue *decimal.Decimal `gorm:"type:decimal(10,2)"` // optional declared value
	IsFragile    bool             `gorm:"default:false"`
	Notes        *string          `gorm:"type:text"`

	SenderName    string  `gorm:"not null"`
	SenderPhone   string  `gorm:"not null;index"`
	SenderAddress *string `gorm:"type:text"`

	ReceiverName    string  `gorm:"not null"`
	ReceiverPhone   string  `gorm:"not null;index"` // looked up against ReceiverFraudProfile.Phone at creation time
	ReceiverEmail   *string `gorm:"type:text"`      // optional — for email notifications
	ReceiverAddress string  `gorm:"not null"`       // frozen text snapshot printed on the shipping label

	// ReceiverAddressID optionally links to a structured, reusable Address
	// record (used for zone/pricing lookups and the sender's saved address
	// book). ReceiverAddress above always holds the frozen text actually
	// printed on the label, independent of whether the Address record is
	// later edited or deleted. (Previously this was an untyped
	// "DeliveryAreaID *string" with no association and no backing model at
	// all — a dangling foreign key to nothing.)
	ReceiverAddressID  *string  `gorm:"type:uuid;index"`
	ReceiverAddressRef *Address `gorm:"foreignKey:ReceiverAddressID"`

	// Financial fields use decimal.Decimal, not float64: float64 cannot
	// represent all base-10 amounts exactly, and COD cash reconciliation /
	// revenue splits need exact arithmetic (a float64 rounding drift of a
	// few paisa compounds across thousands of shipments into a real cash
	// mismatch at hub deposit time).
	DeliveryCharge decimal.Decimal `gorm:"type:decimal(10,2);not null"`
	// codAmount = 0 → PREPAID → OTP. codAmount > 0 → COD → Cash handover (merchant only).
	CodAmount   decimal.Decimal `gorm:"type:decimal(10,2);not null;default:0"`
	TotalCharge decimal.Decimal `gorm:"type:decimal(10,2);not null"`

	// Set automatically based on codAmount at creation time.
	ConfirmationMethod ConfirmationMethod `gorm:"type:varchar(20);not null"`

	// Revenue split (populated after delivery).
	SplitZone     *SplitZone       `gorm:"type:varchar(20)"`
	RiderSharePct *decimal.Decimal `gorm:"type:decimal(5,2)"`
	RiderShare    *decimal.Decimal `gorm:"type:decimal(10,2)"`
	SystemShare   *decimal.Decimal `gorm:"type:decimal(10,2)"`
	MerchantNet   *decimal.Decimal `gorm:"type:decimal(10,2)"`

	// Phase 2 fields (included in schema now so migrations are non-breaking).
	PickupType        PickupType       `gorm:"type:varchar(20);default:'HUB_DROP_OFF'"`
	PickupSurcharge   *decimal.Decimal `gorm:"type:decimal(10,2)"`
	ReturnChargeTotal *decimal.Decimal `gorm:"type:decimal(10,2)"`

	DeliveredAt *time.Time // nullable — only set when status = DELIVERED
	CreatedAt   time.Time  `gorm:"index"` // date-range reporting (daily/weekly volume, SLA dashboards)
	UpdatedAt   time.Time

	// Associations
	Merchant               *Merchant                `gorm:"foreignKey:MerchantID"`
	GuestSender            *GuestSender             `gorm:"foreignKey:GuestSenderID"`
	CreatedByAdmin         *Admin                   `gorm:"foreignKey:CreatedByAdminID"`
	Rider                  *Rider                   `gorm:"foreignKey:RiderID"`
	Hub                    *Hub                     `gorm:"foreignKey:HubID"`
	WalkInPayment          *WalkInPayment           `gorm:"foreignKey:ShipmentID"`
	DeliveryConfirmation   *DeliveryConfirmation    `gorm:"foreignKey:ShipmentID"`
	CodConfirmation        *CodDeliveryConfirmation `gorm:"foreignKey:ShipmentID"`
	DeliveryRating         *DeliveryRating          `gorm:"foreignKey:ShipmentID"`
	MerchantDeliveryRating *MerchantDeliveryRating  `gorm:"foreignKey:ShipmentID"`
}

func (s *Shipment) BeforeCreate(tx *gorm.DB) error {
	if s.ID == "" {
		s.ID = uuid.NewString()
	}
	if err := s.validateSenderReference(); err != nil {
		return err
	}
	return s.deriveConfirmationMethod()
}

func (s *Shipment) BeforeSave(tx *gorm.DB) error {
	return s.validateSenderReference()
}

// validateSenderReference ensures SenderType and the two nullable sender FKs
// agree: a MERCHANT shipment must reference a Merchant (and not a
// GuestSender), and a GUEST shipment the reverse. Previously nothing enforced
// this — a Shipment could be created with SenderType=MERCHANT and no
// MerchantID at all, or with both MerchantID and GuestSenderID set.
func (s *Shipment) validateSenderReference() error {
	switch s.SenderType {
	case SenderTypeMerchant:
		if s.MerchantID == nil || *s.MerchantID == "" {
			return fmt.Errorf("sender type MERCHANT requires MerchantID")
		}
		if s.GuestSenderID != nil {
			return fmt.Errorf("sender type MERCHANT must not set GuestSenderID")
		}
	case SenderTypeGuest:
		if s.GuestSenderID == nil || *s.GuestSenderID == "" {
			return fmt.Errorf("sender type GUEST requires GuestSenderID")
		}
		if s.MerchantID != nil {
			return fmt.Errorf("sender type GUEST must not set MerchantID")
		}
	default:
		return fmt.Errorf("unknown sender type: %q", s.SenderType)
	}
	return nil
}

// deriveConfirmationMethod sets ConfirmationMethod from CodAmount:
// codAmount = 0 → PREPAID → OTP. codAmount > 0 → COD → Cash handover (merchant only).
func (s *Shipment) deriveConfirmationMethod() error {
	if s.CodAmount.IsPositive() {
		if s.SenderType != SenderTypeMerchant {
			return fmt.Errorf("COD shipments (codAmount > 0) are only allowed for merchant senders")
		}
		s.ConfirmationMethod = ConfirmationMethodCashHandover
		return nil
	}
	s.ConfirmationMethod = ConfirmationMethodOTP
	return nil
}
