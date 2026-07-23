package admin

import (
	"errors"
	"testing"
	"time"

	"github.com/gofiber/fiber/v3"
	"server/internal/config"
	"server/internal/domain/admin/dto"
	"server/internal/models"
	querybuilder "server/pkg/queryBuilder"
)

// MockRepository is a manual mock implementation of Repository
type MockRepository struct {
	GetUserByEmailFunc func(email string) (*models.User, error)
	GetHubByIDFunc     func(hubID string) (*models.Hub, error)
	CreateAdminFunc    func(user *models.User, account *models.Account, admin *models.Admin) error
	FindAllFunc        func(c fiber.Ctx, params querybuilder.QueryParams) ([]models.Admin, int64, error)
	FindByIDFunc       func(id string) (*models.Admin, error)
	FindByUserIDFunc   func(userID string) (*models.Admin, error)
	UpdateFunc         func(admin *models.Admin, user *models.User) error
	DeleteFunc         func(id string) error
}

func (m *MockRepository) CreateAdmin(user *models.User, account *models.Account, admin *models.Admin) error {
	if m.CreateAdminFunc != nil {
		return m.CreateAdminFunc(user, account, admin)
	}
	return nil
}

func (m *MockRepository) FindAll(c fiber.Ctx, params querybuilder.QueryParams) ([]models.Admin, int64, error) {
	if m.FindAllFunc != nil {
		return m.FindAllFunc(c, params)
	}
	return nil, 0, nil
}

func (m *MockRepository) FindByID(id string) (*models.Admin, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(id)
	}
	return nil, nil
}

func (m *MockRepository) FindByUserID(userID string) (*models.Admin, error) {
	if m.FindByUserIDFunc != nil {
		return m.FindByUserIDFunc(userID)
	}
	return nil, nil
}

func (m *MockRepository) Update(admin *models.Admin, user *models.User) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(admin, user)
	}
	return nil
}

func (m *MockRepository) Delete(id string) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(id)
	}
	return nil
}

func (m *MockRepository) GetUserByEmail(email string) (*models.User, error) {
	if m.GetUserByEmailFunc != nil {
		return m.GetUserByEmailFunc(email)
	}
	return nil, errors.New("not found")
}

func (m *MockRepository) GetHubByID(hubID string) (*models.Hub, error) {
	if m.GetHubByIDFunc != nil {
		return m.GetHubByIDFunc(hubID)
	}
	return nil, nil
}

// MockEmailSender is a manual mock implementation of email.Sender
type MockEmailSender struct {
	SendEmailFunc              func(to string, subject string, templateFile string, data any) error
	SendVerificationEmailFunc  func(to, name, otp string) error
	SendPasswordResetEmailFunc func(to, name, otp string) error
	SendAdminWelcomeEmailFunc  func(to, name, tempPassword, loginLink string) error
}

func (m *MockEmailSender) SendEmail(to string, subject string, templateFile string, data any) error {
	if m.SendEmailFunc != nil {
		return m.SendEmailFunc(to, subject, templateFile, data)
	}
	return nil
}

func (m *MockEmailSender) SendVerificationEmail(to, name, otp string) error {
	if m.SendVerificationEmailFunc != nil {
		return m.SendVerificationEmailFunc(to, name, otp)
	}
	return nil
}

func (m *MockEmailSender) SendPasswordResetEmail(to, name, otp string) error {
	if m.SendPasswordResetEmailFunc != nil {
		return m.SendPasswordResetEmailFunc(to, name, otp)
	}
	return nil
}

func (m *MockEmailSender) SendAdminWelcomeEmail(to, name, tempPassword, loginLink string) error {
	if m.SendAdminWelcomeEmailFunc != nil {
		return m.SendAdminWelcomeEmailFunc(to, name, tempPassword, loginLink)
	}
	return nil
}

func TestCreateAdmin_Success(t *testing.T) {
	mockRepo := &MockRepository{
		GetUserByEmailFunc: func(email string) (*models.User, error) {
			return nil, errors.New("user not found") // email is free to register
		},
		GetHubByIDFunc: func(hubID string) (*models.Hub, error) {
			return &models.Hub{ID: hubID, Name: "Test Hub"}, nil
		},
		CreateAdminFunc: func(user *models.User, account *models.Account, admin *models.Admin) error {
			// Assertions during creation
			if user.Role != models.RoleAdmin {
				t.Errorf("Expected role ADMIN, got %s", user.Role)
			}
			if user.Status != models.UserStatusActive {
				t.Errorf("Expected status ACTIVE, got %s", user.Status)
			}
			if !user.NeedsPasswordChange {
				t.Errorf("Expected NeedsPasswordChange to be true")
			}
			if admin.Name != "John Doe" {
				t.Errorf("Expected name John Doe, got %s", admin.Name)
			}
			return nil
		},
	}

	welcomeEmailSent := false
	var sentTempPass string
	mockEmail := &MockEmailSender{
		SendAdminWelcomeEmailFunc: func(to, name, tempPassword, loginLink string) error {
			welcomeEmailSent = true
			sentTempPass = tempPassword
			if to != "john@example.com" {
				t.Errorf("Expected email to john@example.com, got %s", to)
			}
			if name != "John Doe" {
				t.Errorf("Expected name John Doe, got %s", name)
			}
			if loginLink != "http://localhost:3000/login" {
				t.Errorf("Expected login link http://localhost:3000/login, got %s", loginLink)
			}
			return nil
		},
	}

	cfg := &config.Config{
		FrontendUrl: "http://localhost:3000",
	}

	svc := NewService(mockRepo, mockEmail, cfg)

	hubID := "12345678-1234-1234-1234-123456789012"
	req := &adminDto.CreateAdminRequest{
		Email:             "john@example.com",
		Name:              "John Doe",
		Phone:             "01711111111",
		District:          models.DistrictDhaka,
		HubID:             &hubID,
		CanApproveCashout: nil,
	}

	res, err := svc.CreateAdmin(req)
	if err != nil {
		t.Fatalf("CreateAdmin failed: %v", err)
	}

	if res.Email != "john@example.com" {
		t.Errorf("Expected response email to be john@example.com, got %s", res.Email)
	}
	if !res.CanApproveCashout {
		t.Errorf("Expected CanApproveCashout default to be true")
	}

	// Verify asynchronous email send
	// We wait slightly since it runs in a goroutine
	for i := 0; i < 50; i++ {
		if welcomeEmailSent {
			break
		}
		time.Sleep(2 * time.Millisecond)
	}

	if !welcomeEmailSent {
		t.Fatalf("Expected admin welcome email to be sent")
	}

	// Verify password generation is populated and length is correct
	if len(sentTempPass) != 12 {
		t.Errorf("Expected generated password of length 12, got %d", len(sentTempPass))
	}
}

func TestCreateAdmin_EmailAlreadyExists(t *testing.T) {
	mockRepo := &MockRepository{
		GetUserByEmailFunc: func(email string) (*models.User, error) {
			return &models.User{ID: "existing-user-id"}, nil
		},
	}
	mockEmail := &MockEmailSender{}
	cfg := &config.Config{}

	svc := NewService(mockRepo, mockEmail, cfg)

	req := &adminDto.CreateAdminRequest{
		Email: "existing@example.com",
		Name:  "Test User",
	}

	_, err := svc.CreateAdmin(req)
	if err == nil || err.Error() != "email is already registered" {
		t.Errorf("Expected 'email is already registered' error, got %v", err)
	}
}
