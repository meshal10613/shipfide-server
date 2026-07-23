package authDto

type RegisterRequest struct {
	Name     string  `json:"name" validate:"required,min=3"`
	Email    string  `json:"email" validate:"required,email"`
	Phone    *string `json:"phone" validate:"omitempty"`
	Password string  `json:"password" validate:"required,password"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type VerifyOtpRequest struct {
	Email string `json:"email" validate:"required,email"`
	Otp   string `json:"otp" validate:"required,len=6"`
}

type ResetPasswordRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Otp      string `json:"otp" validate:"required,len=6"`
	Password string `json:"password" validate:"required,password"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required,password"`
}

type VerifyEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
	Otp   string `json:"otp" validate:"required,len=6"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}
