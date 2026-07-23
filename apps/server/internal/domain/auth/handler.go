package auth

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"

	"server/internal/domain/auth/dto"
	httpresponse "server/pkg/httpResponse"
	"server/pkg/utils"
	"server/pkg/validation"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterUser(c fiber.Ctx) error {
	var req authDto.RegisterRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "invalid request body",
			Details: err.Error(),
		})
	}

	valVal := c.Locals("validator")
	if validatorInstance, ok := valVal.(*validation.CustomValidator); ok {
		if err := validatorInstance.Validate(&req); err != nil {
			return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
				Success: false,
				Message: "validation failed",
				Details: err.Error(),
			})
		}
	}

	res, err := h.service.Register(&req, c.UserAgent(), c.IP())
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	// Set cookies for access, refresh, and session tokens
	utils.SetAuthCookies(c, res.AccessToken, res.RefreshToken, res.SessionToken)

	return c.Status(http.StatusCreated).JSON(httpresponse.Success{
		Success: true,
		Message: "registration successful",
		Data:    res,
	})
}

func (h *Handler) LoginUser(c fiber.Ctx) error {
	var req authDto.LoginRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "invalid request body",
			Details: err.Error(),
		})
	}

	valVal := c.Locals("validator")
	if validatorInstance, ok := valVal.(*validation.CustomValidator); ok {
		if err := validatorInstance.Validate(&req); err != nil {
			return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
				Success: false,
				Message: "validation failed",
				Details: err.Error(),
			})
		}
	}

	res, err := h.service.Login(&req, c.UserAgent(), c.IP())
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	// Set cookies for access, refresh, and session tokens
	utils.SetAuthCookies(c, res.AccessToken, res.RefreshToken, res.SessionToken)

	return c.Status(http.StatusOK).JSON(httpresponse.Success{
		Success: true,
		Message: "login successful",
		Data:    res,
	})
}

func (h *Handler) GetMe(c fiber.Ctx) error {
	userIDVal := c.Locals("user_id")
	var userIDStr string

	if u, ok := userIDVal.(uuid.UUID); ok {
		userIDStr = u.String()
	} else if s, ok := userIDVal.(string); ok {
		userIDStr = s
	}

	user, err := h.service.GetMe(userIDStr)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(httpresponse.Error{
			Success: false,
			Message: "failed to retrieve profile",
		})
	}

	return c.Status(http.StatusOK).JSON(httpresponse.Success{
		Success: true,
		Message: "profile retrieved successfully",
		Data:    user,
	})
}

func (h *Handler) LogoutUser(c fiber.Ctx) error {
	sessionToken := c.Cookies("session_token")
	if sessionToken == "" {
		// fallback to parsing from cookie or header if not found
		sessionToken = c.Get("X-Session-Token")
	}

	if sessionToken != "" {
		_ = h.service.Logout(sessionToken)
	}

	utils.ClearAuthCookies(c)

	return c.Status(http.StatusOK).JSON(httpresponse.Success{
		Success: true,
		Message: "logout successful",
	})
}

func (h *Handler) RefreshToken(c fiber.Ctx) error {
	// 1. Get refresh token from cookie or request body fallback
	refreshToken := c.Cookies("refresh_token")
	var req authDto.RefreshTokenRequest

	if refreshToken == "" {
		if err := c.Bind().Body(&req); err == nil {
			refreshToken = req.RefreshToken
		}
	}

	if refreshToken == "" {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "refresh token is required",
		})
	}

	sessionToken := c.Cookies("session_token")

	// 2. Call service to refresh
	res, err := h.service.RefreshToken(refreshToken, sessionToken, c.UserAgent(), c.IP())
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	// 3. Set new cookies
	utils.SetAuthCookies(c, res.AccessToken, res.RefreshToken, res.SessionToken)

	return c.Status(http.StatusOK).JSON(httpresponse.Success{
		Success: true,
		Message: "token refreshed successfully",
		Data:    res,
	})
}

func (h *Handler) ForgotPassword(c fiber.Ctx) error {
	var req authDto.ForgotPasswordRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "invalid request body",
			Details: err.Error(),
		})
	}

	valVal := c.Locals("validator")
	if validatorInstance, ok := valVal.(*validation.CustomValidator); ok {
		if err := validatorInstance.Validate(&req); err != nil {
			return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
				Success: false,
				Message: "validation failed",
				Details: err.Error(),
			})
		}
	}

	if err := h.service.ForgotPassword(req.Email); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(httpresponse.Success{
		Success: true,
		Message: "If your email is registered, a password reset OTP code has been sent.",
	})
}

func (h *Handler) VerifyOtp(c fiber.Ctx) error {
	var req authDto.VerifyOtpRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "invalid request body",
			Details: err.Error(),
		})
	}

	valVal := c.Locals("validator")
	if validatorInstance, ok := valVal.(*validation.CustomValidator); ok {
		if err := validatorInstance.Validate(&req); err != nil {
			return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
				Success: false,
				Message: "validation failed",
				Details: err.Error(),
			})
		}
	}

	if err := h.service.VerifyOtp(req.Email, req.Otp); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(httpresponse.Success{
		Success: true,
		Message: "OTP code verified successfully.",
	})
}

func (h *Handler) ResetPassword(c fiber.Ctx) error {
	var req authDto.ResetPasswordRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "invalid request body",
			Details: err.Error(),
		})
	}

	valVal := c.Locals("validator")
	if validatorInstance, ok := valVal.(*validation.CustomValidator); ok {
		if err := validatorInstance.Validate(&req); err != nil {
			return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
				Success: false,
				Message: "validation failed",
				Details: err.Error(),
			})
		}
	}

	if err := h.service.ResetPassword(req.Email, req.Otp, req.Password); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(httpresponse.Success{
		Success: true,
		Message: "Password has been reset successfully.",
	})
}

func (h *Handler) ChangePassword(c fiber.Ctx) error {
	var req authDto.ChangePasswordRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "invalid request body",
			Details: err.Error(),
		})
	}

	valVal := c.Locals("validator")
	if validatorInstance, ok := valVal.(*validation.CustomValidator); ok {
		if err := validatorInstance.Validate(&req); err != nil {
			return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
				Success: false,
				Message: "validation failed",
				Details: err.Error(),
			})
		}
	}

	userIDVal := c.Locals("user_id")
	var userIDStr string
	if u, ok := userIDVal.(uuid.UUID); ok {
		userIDStr = u.String()
	} else if s, ok := userIDVal.(string); ok {
		userIDStr = s
	}

	if err := h.service.ChangePassword(userIDStr, req.OldPassword, req.NewPassword); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(httpresponse.Success{
		Success: true,
		Message: "Password changed successfully.",
	})
}

func (h *Handler) SendVerificationOtp(c fiber.Ctx) error {
	var req struct {
		Email string `json:"email"`
	}

	// First try to bind body to see if email was provided
	_ = c.Bind().Body(&req)

	emailStr := req.Email
	if emailStr == "" {
		// fallback to local context email if user is authenticated
		emailVal := c.Locals("email")
		if e, ok := emailVal.(string); ok {
			emailStr = e
		}
	}

	if emailStr == "" {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "email is required",
		})
	}

	if err := h.service.SendVerificationOtp(emailStr); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(httpresponse.Success{
		Success: true,
		Message: "Verification OTP code sent successfully.",
	})
}

func (h *Handler) VerifyEmail(c fiber.Ctx) error {
	var req authDto.VerifyEmailRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "invalid request body",
			Details: err.Error(),
		})
	}

	valVal := c.Locals("validator")
	if validatorInstance, ok := valVal.(*validation.CustomValidator); ok {
		if err := validatorInstance.Validate(&req); err != nil {
			return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
				Success: false,
				Message: "validation failed",
				Details: err.Error(),
			})
		}
	}

	if err := h.service.VerifyEmail(req.Email, req.Otp); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(httpresponse.Success{
		Success: true,
		Message: "Email address verified successfully.",
	})
}
