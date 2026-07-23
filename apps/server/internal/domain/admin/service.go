package admin

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"
	"server/internal/config"
	"server/internal/domain/admin/dto"
	"server/internal/models"
	"server/pkg/email"
	querybuilder "server/pkg/queryBuilder"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type Service interface {
	CreateAdmin(req *adminDto.CreateAdminRequest) (*adminDto.AdminResponse, error)
	GetAdmins(c fiber.Ctx, params querybuilder.QueryParams) ([]*adminDto.AdminResponse, int64, error)
	GetAdmin(id string) (*adminDto.AdminResponse, error)
	UpdateAdmin(id string, req *adminDto.UpdateAdminRequest) (*adminDto.AdminResponse, error)
	DeleteAdmin(id string) error
}

type service struct {
	repo        Repository
	emailSender email.Sender
	config      *config.Config
}

func NewService(repo Repository, emailSender email.Sender, cfg *config.Config) Service {
	return &service{
		repo:        repo,
		emailSender: emailSender,
		config:      cfg,
	}
}

func generateTempPassword(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*"
	password := make([]byte, length)
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		password[i] = charset[num.Int64()]
	}
	return string(password), nil
}

func (s *service) CreateAdmin(req *adminDto.CreateAdminRequest) (*adminDto.AdminResponse, error) {
	// 1. Verify email uniqueness
	_, err := s.repo.GetUserByEmail(req.Email)
	if err == nil {
		return nil, errors.New("email is already registered")
	}

	// 2. Validate Hub existence if HubID is provided
	if req.HubID != nil && *req.HubID != "" {
		r, err := s.repo.GetHubByID(*req.HubID)
		if err != nil || r == nil {
			return nil, fmt.Errorf("invalid hub: hub not found for ID %s", *req.HubID)
		}
	}

	// 3. Generate secure temporary password
	tempPassword, err := generateTempPassword(12)
	if err != nil {
		return nil, fmt.Errorf("failed to generate temporary password: %w", err)
	}

	// 4. Create models
	userID := uuid.NewString()
	user := &models.User{
		ID:                  userID,
		Name:                req.Name,
		Email:               req.Email,
		Role:                models.RoleAdmin,
		Status:              models.UserStatusActive, // Activated since created by super admin
		NeedsPasswordChange: true,              // Enforce change on first login
	}

	account := &models.Account{
		ID:     uuid.NewString(),
		UserID: userID,
	}
	if err := account.HashPassword(tempPassword); err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	admin := &models.Admin{
		ID:       uuid.NewString(),
		UserID:   userID,
		Name:     req.Name,
		Phone:    req.Phone,
		District: req.District,
		HubID:    req.HubID,
	}

	if req.CanApproveCashout != nil {
		admin.CanApproveCashout = *req.CanApproveCashout
	} else {
		admin.CanApproveCashout = true
	}
	if req.CanManageRiders != nil {
		admin.CanManageRiders = *req.CanManageRiders
	} else {
		admin.CanManageRiders = true
	}
	if req.CanViewFraudFlags != nil {
		admin.CanViewFraudFlags = *req.CanViewFraudFlags
	} else {
		admin.CanViewFraudFlags = true
	}

	// 5. Store in DB (Transaction)
	if err := s.repo.CreateAdmin(user, account, admin); err != nil {
		return nil, err
	}

	// Preload nested fields for the response mapping
	admin.User = user

	// 6. Send Invitation Email Asynchronously
	loginLink := fmt.Sprintf("%s/login", s.config.FrontendUrl)
	go func() {
		if err := s.emailSender.SendAdminWelcomeEmail(user.Email, user.Name, tempPassword, loginLink); err != nil {
			log.Printf("failed to send welcome email to new admin %s: %v", user.Email, err)
		}
	}()

	return adminDto.MapToAdminResponse(admin), nil
}

func (s *service) GetAdmins(c fiber.Ctx, params querybuilder.QueryParams) ([]*adminDto.AdminResponse, int64, error) {
	admins, total, err := s.repo.FindAll(c, params)
	if err != nil {
		return nil, 0, err
	}
	return adminDto.MapToAdminResponseList(admins), total, nil
}

func (s *service) GetAdmin(id string) (*adminDto.AdminResponse, error) {
	admin, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("admin not found")
	}
	return adminDto.MapToAdminResponse(admin), nil
}

func (s *service) UpdateAdmin(id string, req *adminDto.UpdateAdminRequest) (*adminDto.AdminResponse, error) {
	admin, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("admin not found")
	}

	user := admin.User

	// Apply updates to Admin & nested User
	if req.Name != nil {
		admin.Name = *req.Name
		if user != nil {
			user.Name = *req.Name
		}
	}
	if req.Phone != nil {
		admin.Phone = *req.Phone
	}
	if req.District != nil {
		admin.District = *req.District
	}
	if req.HubID != nil {
		// Double pointer parsing to handle nullify vs unchanged vs updated
		admin.HubID = *req.HubID
		if admin.HubID != nil && *admin.HubID != "" {
			// Validate that the hub exists
			_, err := s.repo.GetHubByID(*admin.HubID)
			if err != nil {
				return nil, fmt.Errorf("invalid hub: hub not found for ID %s", *admin.HubID)
			}
		}
	}
	if req.CanApproveCashout != nil {
		admin.CanApproveCashout = *req.CanApproveCashout
	}
	if req.CanManageRiders != nil {
		admin.CanManageRiders = *req.CanManageRiders
	}
	if req.CanViewFraudFlags != nil {
		admin.CanViewFraudFlags = *req.CanViewFraudFlags
	}

	if err := s.repo.Update(admin, user); err != nil {
		return nil, err
	}

	return adminDto.MapToAdminResponse(admin), nil
}

func (s *service) DeleteAdmin(id string) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("admin not found")
	}
	return s.repo.Delete(id)
}
