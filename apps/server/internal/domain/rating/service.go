package rating

import (
	"fmt"
	"time"

	"server/internal/domain/rating/dto"
	"server/internal/models"
)

type Service interface {
	CreateDeliveryRating(req dto.CreateDeliveryRatingRequest) (*dto.DeliveryRatingResponse, error)
	ListDeliveryRatingsByRiderID(riderID string) ([]dto.DeliveryRatingResponse, error)
	CreateMerchantRating(req dto.CreateMerchantRatingRequest) (*dto.MerchantRatingResponse, error)
	ListMerchantRatings() ([]dto.MerchantRatingResponse, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateDeliveryRating(req dto.CreateDeliveryRatingRequest) (*dto.DeliveryRatingResponse, error) {
	now := time.Now()
	rating := &models.DeliveryRating{
		ShipmentID:      req.ShipmentID,
		RiderID:         req.RiderID,
		ReceiverPhone:   req.ReceiverPhone,
		IsAnonymous:     req.IsAnonymous,
		Stars:           req.Stars,
		Comment:         req.Comment,
		RatingWindowEnd: now.Add(7 * 24 * time.Hour),
		SubmittedAt:     &now,
	}

	if err := s.repo.CreateDeliveryRating(rating); err != nil {
		return nil, fmt.Errorf("failed to create delivery rating: %w", err)
	}

	// Update rider rating cache
	var rider models.Rider
	if err := s.repo.DB().First(&rider, "id = ?", req.RiderID).Error; err == nil {
		rider.TotalRatingSum += float64(req.Stars)
		rider.TotalRatingCount++
		if rider.TotalRatingCount > 0 {
			rider.AverageRating = rider.TotalRatingSum / float64(rider.TotalRatingCount)
		}

		if req.Stars <= 2 {
			rider.ConsecutiveLowRatings++
		} else {
			rider.ConsecutiveLowRatings = 0
		}

		switch {
		case rider.TotalRatingCount < 5:
			rider.RatingBadge = models.RiderBadgeNewRider
		case rider.AverageRating >= 4.5:
			rider.RatingBadge = models.RiderBadgeTopRider
		case rider.AverageRating >= 3.5:
			rider.RatingBadge = models.RiderBadgeGood
		case rider.AverageRating >= 2.5:
			rider.RatingBadge = models.RiderBadgeAverage
		default:
			rider.RatingBadge = models.RiderBadgeUnderReview
		}

		_ = s.repo.DB().Save(&rider).Error
	}

	res := dto.ToDeliveryRatingResponse(rating)
	return &res, nil
}

func (s *service) ListDeliveryRatingsByRiderID(riderID string) ([]dto.DeliveryRatingResponse, error) {
	ratings, err := s.repo.ListDeliveryRatingsByRiderID(riderID)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.DeliveryRatingResponse, 0, len(ratings))
	for _, r := range ratings {
		responses = append(responses, dto.ToDeliveryRatingResponse(&r))
	}
	return responses, nil
}

func (s *service) CreateMerchantRating(req dto.CreateMerchantRatingRequest) (*dto.MerchantRatingResponse, error) {
	var merchant models.Merchant
	if err := s.repo.DB().First(&merchant, "id = ?", req.MerchantID).Error; err != nil {
		return nil, fmt.Errorf("merchant not found")
	}

	if !merchant.HasCompletedFirstDelivery {
		return nil, fmt.Errorf("merchant rating is only allowed after completing the first delivery")
	}

	now := time.Now()
	rating := &models.MerchantDeliveryRating{
		ShipmentID:      req.ShipmentID,
		MerchantID:      req.MerchantID,
		Stars:           req.Stars,
		Comment:         req.Comment,
		Tag:             req.Tag,
		RatingWindowEnd: now.Add(14 * 24 * time.Hour),
		SubmittedAt:     &now,
	}

	if err := s.repo.CreateMerchantRating(rating); err != nil {
		return nil, fmt.Errorf("failed to submit merchant delivery rating: %w", err)
	}

	res := dto.ToMerchantRatingResponse(rating)
	return &res, nil
}

func (s *service) ListMerchantRatings() ([]dto.MerchantRatingResponse, error) {
	ratings, err := s.repo.ListMerchantRatings()
	if err != nil {
		return nil, err
	}

	responses := make([]dto.MerchantRatingResponse, 0, len(ratings))
	for _, m := range ratings {
		responses = append(responses, dto.ToMerchantRatingResponse(&m))
	}
	return responses, nil
}
