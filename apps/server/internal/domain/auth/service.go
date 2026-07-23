package auth

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/google/uuid"

	"server/internal/domain/auth/dto"
	"server/internal/models"
	"server/pkg/email"
	"server/pkg/utils"
)

type Service interface {
	Register(req *authDto.RegisterRequest, userAgent, ipAddress string) (*authDto.AuthResponse, error)
	Login(req *authDto.LoginRequest, userAgent, ipAddress string) (*authDto.AuthResponse, error)
	GetMe(userID string) (*models.User, error)
	Logout(sessionToken string) error
	RefreshToken(refreshToken, sessionToken, userAgent, ipAddress string) (*authDto.AuthResponse, error)
	ForgotPassword(email string) error
	VerifyOtp(email, otp string) error
	ResetPassword(email, otp, newPassword string) error
	ChangePassword(userID, oldPassword, newPassword string) error
	SendVerificationOtp(email string) error
	VerifyEmail(email, otp string) error
}

type service struct {
	repo        Repository
	jwt         utils.JwtService
	emailSender email.Sender
}

func NewService(repo Repository, jwt utils.JwtService, emailSender email.Sender) Service {
	return &service{repo: repo, jwt: jwt, emailSender: emailSender}
}

func generateOTP() string {
	n, _ := rand.Int(rand.Reader, big.NewInt(900000))
	return fmt.Sprintf("%06d", n.Int64()+100000)
}

func (s *service) Register(req *authDto.RegisterRequest, userAgent, ipAddress string) (*authDto.AuthResponse, error) {
	// Check if email already exists
	_, err := s.repo.GetUserByEmail(req.Email)
	if err == nil {
		return nil, errors.New("email already registered")
	}

	// Refresh and Session tokens expire in 30 days
	expiresAt := time.Now().Add(30 * 24 * time.Hour)

	// Generate UUID for user
	userID := uuid.NewString()

	// Generate Access Token (JWT)
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	accessToken, err := s.jwt.GenerateToken(userUUID, req.Name, req.Email, string(models.RoleMerchant))
	if err != nil {
		return nil, err
	}

	// Generate Refresh Token & Session Token
	refreshToken := uuid.NewString()
	sessionToken := uuid.NewString()

	otp := generateOTP()
	otpExpiresAt := time.Now().Add(15 * time.Minute)

	user := &models.User{
		ID:     userID,
		Name:   req.Name,
		Email:  req.Email,
		Role:   models.RoleMerchant,      // default role
		Status: models.UserStatusPending, // register as pending email verification
	}

	account := &models.Account{
		ID:                    uuid.NewString(),
		RefreshToken:          &refreshToken,
		RefreshTokenExpiresAt: &expiresAt,
		Otp:                   &otp,
		OtpExpiresAt:          &otpExpiresAt,
	}

	if err := account.HashPassword(req.Password); err != nil {
		return nil, err
	}

	if err := s.repo.CreateUser(user, account); err != nil {
		return nil, err
	}

	// Create Session
	session := &models.Session{
		SessionToken: sessionToken,
		UserID:       user.ID,
		ExpiresAt:    expiresAt,
	}
	if userAgent != "" {
		session.UserAgent = &userAgent
	}
	if ipAddress != "" {
		session.IPAddress = &ipAddress
	}

	if err := s.repo.CreateSession(session); err != nil {
		return nil, err
	}

	// Send verification email asynchronously
	go func() {
		if err := s.emailSender.SendVerificationEmail(user.Email, user.Name, otp); err != nil {
			log.Printf("failed to send verification email: %v", err)
		}
	}()

	preloadedUser, err := s.repo.GetUserByID(user.ID)
	if err != nil {
		return nil, err
	}

	return &authDto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		SessionToken: sessionToken,
		User:         preloadedUser,
	}, nil
}

func (s *service) Login(req *authDto.LoginRequest, userAgent, ipAddress string) (*authDto.AuthResponse, error) {
	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !user.Account.CheckPassword(req.Password) {
		return nil, errors.New("invalid email or password")
	}

	// Check active sessions limit (max 3). Auto-delete all sessions if limit is reached.
	activeCount, err := s.repo.GetActiveSessionsCount(user.ID)
	if err == nil && activeCount >= 3 {
		_ = s.repo.DeleteAllSessionsByUserID(user.ID)
	}

	// Refresh and Session tokens expire in 30 days
	expiresAt := time.Now().Add(30 * 24 * time.Hour)

	// Generate Access Token (JWT)
	userUUID, err := uuid.Parse(user.ID)
	if err != nil {
		return nil, err
	}

	accessToken, err := s.jwt.GenerateToken(userUUID, user.Name, user.Email, string(user.Role))
	if err != nil {
		return nil, err
	}

	// Generate Refresh Token & Session Token
	refreshToken := uuid.NewString()
	sessionToken := uuid.NewString()

	// Update Account Refresh Token details
	user.Account.RefreshToken = &refreshToken
	user.Account.RefreshTokenExpiresAt = &expiresAt
	if err := s.repo.UpdateAccount(user.Account); err != nil {
		return nil, err
	}

	// Create Session
	session := &models.Session{
		SessionToken: sessionToken,
		UserID:       user.ID,
		ExpiresAt:    expiresAt,
	}
	if userAgent != "" {
		session.UserAgent = &userAgent
	}
	if ipAddress != "" {
		session.IPAddress = &ipAddress
	}

	if err := s.repo.CreateSession(session); err != nil {
		return nil, err
	}

	preloadedUser, err := s.repo.GetUserByID(user.ID)
	if err != nil {
		return nil, err
	}

	return &authDto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		SessionToken: sessionToken,
		User:         preloadedUser,
	}, nil
}

func (s *service) GetMe(userID string) (*models.User, error) {
	return s.repo.GetUserByIDWithSessions(userID)
}

func (s *service) Logout(sessionToken string) error {
	if sessionToken == "" {
		return errors.New("invalid session token")
	}
	session, err := s.repo.GetSessionByToken(sessionToken)
	if err != nil {
		return nil // already logged out or session not found
	}

	now := time.Now()
	session.RevokedAt = &now
	if err := s.repo.UpdateSession(session); err != nil {
		return err
	}

	// Also clear refresh token from account associated with this session user
	account, err := s.repo.GetAccountByUserID(session.UserID)
	if err == nil && account != nil {
		account.RefreshToken = nil
		account.RefreshTokenExpiresAt = nil
		_ = s.repo.UpdateAccount(account)
	}

	return s.repo.DeleteSessionByToken(sessionToken)
}

func (s *service) RefreshToken(refreshToken, sessionToken, userAgent, ipAddress string) (*authDto.AuthResponse, error) {
	account, err := s.repo.GetAccountByRefreshToken(refreshToken)
	if err != nil {
		return nil, errors.New("invalid or expired refresh token")
	}

	if account.RefreshTokenExpiresAt == nil || account.RefreshTokenExpiresAt.Before(time.Now()) {
		return nil, errors.New("refresh token expired")
	}

	user, err := s.repo.GetUserByID(account.UserID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Generate new access, refresh, and session tokens
	expiresAt := time.Now().Add(30 * 24 * time.Hour)
	userUUID, err := uuid.Parse(user.ID)
	if err != nil {
		return nil, err
	}

	newAccessToken, err := s.jwt.GenerateToken(userUUID, user.Name, user.Email, string(user.Role))
	if err != nil {
		return nil, err
	}

	newRefreshToken := uuid.NewString()
	newSessionToken := uuid.NewString()

	// Update refresh token details
	account.RefreshToken = &newRefreshToken
	account.RefreshTokenExpiresAt = &expiresAt
	if err := s.repo.UpdateAccount(account); err != nil {
		return nil, err
	}

	// Revoke old session if provided
	if sessionToken != "" {
		_ = s.Logout(sessionToken)
	}

	// Create new Session
	session := &models.Session{
		SessionToken: newSessionToken,
		UserID:       user.ID,
		ExpiresAt:    expiresAt,
	}
	if userAgent != "" {
		session.UserAgent = &userAgent
	}
	if ipAddress != "" {
		session.IPAddress = &ipAddress
	}

	if err := s.repo.CreateSession(session); err != nil {
		return nil, err
	}

	preloadedUser, err := s.repo.GetUserByID(user.ID)
	if err != nil {
		return nil, err
	}

	return &authDto.AuthResponse{
		User:         preloadedUser,
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
		SessionToken: newSessionToken,
	}, nil
}

func (s *service) ForgotPassword(email string) error {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		// Return generic success to prevent user enumeration
		return nil
	}

	otp := generateOTP()
	expiry := time.Now().Add(15 * time.Minute)

	user.Account.Otp = &otp
	user.Account.OtpExpiresAt = &expiry

	if err := s.repo.UpdateAccount(user.Account); err != nil {
		return err
	}

	// Send password reset email asynchronously
	go func() {
		if err := s.emailSender.SendPasswordResetEmail(user.Email, user.Name, otp); err != nil {
			log.Printf("failed to send password reset email: %v", err)
		}
	}()

	return nil
}

func (s *service) VerifyOtp(email, otp string) error {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return errors.New("invalid email or verification code")
	}

	if user.Account.Otp == nil || *user.Account.Otp != otp {
		return errors.New("invalid verification code")
	}

	if user.Account.OtpExpiresAt == nil || user.Account.OtpExpiresAt.Before(time.Now()) {
		return errors.New("verification code expired")
	}

	return nil
}

func (s *service) ResetPassword(email, otp, newPassword string) error {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return errors.New("invalid email or verification code")
	}

	if err := s.VerifyOtp(email, otp); err != nil {
		return err
	}

	// Clear OTP and set new password
	user.Account.Otp = nil
	user.Account.OtpExpiresAt = nil
	if err := user.Account.HashPassword(newPassword); err != nil {
		return err
	}

	user.NeedsPasswordChange = false
	if err := s.repo.UpdateUser(user); err != nil {
		return err
	}

	return s.repo.UpdateAccount(user.Account)
}

func (s *service) ChangePassword(userID, oldPassword, newPassword string) error {
	account, err := s.repo.GetAccountByUserID(userID)
	if err != nil {
		return errors.New("account not found")
	}

	if !account.CheckPassword(oldPassword) {
		return errors.New("incorrect old password")
	}

	if err := account.HashPassword(newPassword); err != nil {
		return err
	}

	user, err := s.repo.GetUserByID(userID)
	if err == nil && user != nil {
		user.NeedsPasswordChange = false
		if err := s.repo.UpdateUser(user); err != nil {
			return err
		}
	}

	return s.repo.UpdateAccount(account)
}

func (s *service) SendVerificationOtp(email string) error {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return errors.New("user not found")
	}

	if user.Status == models.UserStatusActive {
		return errors.New("email is already verified")
	}

	otp := generateOTP()
	expiry := time.Now().Add(15 * time.Minute)

	user.Account.Otp = &otp
	user.Account.OtpExpiresAt = &expiry

	if err := s.repo.UpdateAccount(user.Account); err != nil {
		return err
	}

	// Send verification email asynchronously
	go func() {
		if err := s.emailSender.SendVerificationEmail(user.Email, user.Name, otp); err != nil {
			log.Printf("failed to send verification email: %v", err)
		}
	}()

	return nil
}

func (s *service) VerifyEmail(email, otp string) error {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return errors.New("user not found")
	}

	if user.Status == models.UserStatusActive {
		return errors.New("email is already verified")
	}

	if err := s.VerifyOtp(email, otp); err != nil {
		return err
	}

	// Activate user and clear OTP
	user.Status = models.UserStatusActive
	user.Account.Otp = nil
	user.Account.OtpExpiresAt = nil

	if err := s.repo.UpdateUser(user); err != nil {
		return err
	}

	return s.repo.UpdateAccount(user.Account)
}
