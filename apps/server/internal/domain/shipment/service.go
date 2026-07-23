package shipment

import (
	"fmt"
	"time"

	"server/internal/domain/shipment/dto"
	"server/internal/models"
	"server/pkg/pricing"
	"server/pkg/utils"

	"gorm.io/gorm"
)

type Service interface {
	CalculatePrice(req dto.CalculatePriceRequest) (pricing.PricingResult, error)
	CreateShipment(creatorAdminID *string, req dto.CreateShipmentRequest) (*dto.ShipmentResponse, error)
	GetShipmentByID(id string) (*dto.ShipmentResponse, error)
	GetShipmentByTrackingCode(trackingCode string) (*dto.ShipmentResponse, error)
	UpdateStatus(id string, req dto.UpdateShipmentStatusRequest) (*dto.ShipmentResponse, error)
	AssignRider(id string, riderID string) (*dto.ShipmentResponse, error)
	AssignHub(id string, hubID string) (*dto.ShipmentResponse, error)
	ListShipments(filter ListFilter) ([]dto.ShipmentResponse, int64, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CalculatePrice(req dto.CalculatePriceRequest) (pricing.PricingResult, error) {
	pricingReq := pricing.PricingRequest{
		Zone:            req.Zone,
		WeightGrams:     req.WeightGrams,
		ParcelType:      req.ParcelType,
		IsFragile:       req.IsFragile,
		PickupSurcharge: req.PickupSurcharge,
	}
	return pricing.CalculateDeliveryCharge(pricingReq), nil
}

func (s *service) CreateShipment(creatorAdminID *string, req dto.CreateShipmentRequest) (*dto.ShipmentResponse, error) {
	// 1. Receiver fraud check
	var fraudProfile models.ReceiverFraudProfile
	err := s.repo.DB().First(&fraudProfile, "phone = ?", req.ReceiverPhone).Error
	if err == nil {
		if fraudProfile.CODBlocked && req.CodAmount.IsPositive() {
			return nil, fmt.Errorf("COD is blocked for receiver phone %s due to high fraud risk", req.ReceiverPhone)
		}
	} else if err == gorm.ErrRecordNotFound {
		// Create new profile
		newProfile := models.ReceiverFraudProfile{
			Name:  req.ReceiverName,
			Phone: req.ReceiverPhone,
			Email: req.ReceiverEmail,
		}
		_ = s.repo.DB().Create(&newProfile).Error
	}

	// 2. Price calculation
	priceRes := pricing.CalculateDeliveryCharge(pricing.PricingRequest{
		Zone:        req.ZoneType,
		WeightGrams: req.Weight,
		ParcelType:  req.ParcelType,
		IsFragile:   req.IsFragile,
	})

	// 3. Unique tracking code
	trackingCode := utils.GenerateTrackingCode()

	confirmationMethod := models.ConfirmationMethodOTP
	if req.CodAmount.IsPositive() {
		confirmationMethod = models.ConfirmationMethodCashHandover
	}

	shipment := &models.Shipment{
		TrackingCode:       trackingCode,
		SenderType:         req.SenderType,
		MerchantID:         req.MerchantID,
		GuestSenderID:      req.GuestSenderID,
		CreatedByAdminID:   creatorAdminID,
		HubID:              req.HubID,
		Status:             models.ShipmentPending,
		ParcelType:         req.ParcelType,
		Weight:             req.Weight,
		Dimensions:         req.Dimensions,
		ProductName:        req.ProductName,
		ProductValue:       req.ProductValue,
		IsFragile:          req.IsFragile,
		Notes:              req.Notes,
		SenderName:         req.SenderName,
		SenderPhone:        req.SenderPhone,
		SenderAddress:      req.SenderAddress,
		ReceiverName:       req.ReceiverName,
		ReceiverPhone:      req.ReceiverPhone,
		ReceiverEmail:      req.ReceiverEmail,
		ReceiverAddress:    req.ReceiverAddress,
		ReceiverAddressID:  req.ReceiverAddressID,
		DeliveryCharge:     priceRes.DeliveryCharge,
		CodAmount:          req.CodAmount,
		TotalCharge:        priceRes.TotalCharge,
		ConfirmationMethod: confirmationMethod,
	}

	if err := s.repo.Create(shipment); err != nil {
		return nil, err
	}

	res := dto.ToShipmentResponse(shipment)
	return &res, nil
}

func (s *service) GetShipmentByID(id string) (*dto.ShipmentResponse, error) {
	shipment, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if shipment == nil {
		return nil, fmt.Errorf("shipment not found")
	}

	res := dto.ToShipmentResponse(shipment)
	return &res, nil
}

func (s *service) GetShipmentByTrackingCode(trackingCode string) (*dto.ShipmentResponse, error) {
	shipment, err := s.repo.GetByTrackingCode(trackingCode)
	if err != nil {
		return nil, err
	}
	if shipment == nil {
		return nil, fmt.Errorf("shipment not found")
	}

	res := dto.ToShipmentResponse(shipment)
	return &res, nil
}

func (s *service) UpdateStatus(id string, req dto.UpdateShipmentStatusRequest) (*dto.ShipmentResponse, error) {
	shipment, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if shipment == nil {
		return nil, fmt.Errorf("shipment not found")
	}

	shipment.Status = req.Status
	if req.Notes != nil {
		shipment.Notes = req.Notes
	}

	// Trigger delivery completions & revenue splits
	if req.Status == models.ShipmentDelivered {
		now := time.Now()
		shipment.DeliveredAt = &now

		splitZone := models.SplitSameCity
		splitRes := pricing.CalculateRevenueSplit(splitZone, shipment.DeliveryCharge, shipment.CodAmount)

		shipment.SplitZone = &splitZone
		shipment.RiderSharePct = &splitRes.RiderSharePct
		shipment.RiderShare = &splitRes.RiderShare
		shipment.SystemShare = &splitRes.SystemShare
		shipment.MerchantNet = &splitRes.MerchantNet

		// Update merchant first delivery status
		if shipment.MerchantID != nil {
			var m models.Merchant
			if err := s.repo.DB().First(&m, "id = ?", *shipment.MerchantID).Error; err == nil {
				if !m.HasCompletedFirstDelivery {
					m.HasCompletedFirstDelivery = true
					_ = s.repo.DB().Save(&m).Error
				}
			}
		}

		// Update receiver fraud profile delivery statistics
		var fraudProfile models.ReceiverFraudProfile
		if err := s.repo.DB().First(&fraudProfile, "phone = ?", shipment.ReceiverPhone).Error; err == nil {
			fraudProfile.TotalOrders++
			fraudProfile.TotalDelivered++
			fraudProfile.RefreshScore()
			_ = s.repo.DB().Save(&fraudProfile).Error
		}
	} else if req.Status == models.ShipmentFailedDelivery {
		var fraudProfile models.ReceiverFraudProfile
		if err := s.repo.DB().First(&fraudProfile, "phone = ?", shipment.ReceiverPhone).Error; err == nil {
			fraudProfile.TotalOrders++
			fraudProfile.RefreshScore()
			_ = s.repo.DB().Save(&fraudProfile).Error
		}
	}

	if err := s.repo.Update(shipment); err != nil {
		return nil, err
	}

	res := dto.ToShipmentResponse(shipment)
	return &res, nil
}

func (s *service) AssignRider(id string, riderID string) (*dto.ShipmentResponse, error) {
	shipment, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if shipment == nil {
		return nil, fmt.Errorf("shipment not found")
	}

	shipment.RiderID = &riderID
	if shipment.Status == models.ShipmentPending || shipment.Status == models.ShipmentPickupRequested {
		shipment.Status = models.ShipmentPickupRiderAssigned
	}

	if err := s.repo.Update(shipment); err != nil {
		return nil, err
	}

	res := dto.ToShipmentResponse(shipment)
	return &res, nil
}

func (s *service) AssignHub(id string, hubID string) (*dto.ShipmentResponse, error) {
	shipment, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if shipment == nil {
		return nil, fmt.Errorf("shipment not found")
	}

	shipment.HubID = &hubID
	if err := s.repo.Update(shipment); err != nil {
		return nil, err
	}

	res := dto.ToShipmentResponse(shipment)
	return &res, nil
}

func (s *service) ListShipments(filter ListFilter) ([]dto.ShipmentResponse, int64, error) {
	shipments, total, err := s.repo.List(filter)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]dto.ShipmentResponse, 0, len(shipments))
	for _, item := range shipments {
		responses = append(responses, dto.ToShipmentResponse(&item))
	}
	return responses, total, nil
}
