package auth

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"server/internal/config"
	"server/pkg/email"
	"server/pkg/middlewares"
	"server/pkg/utils"
	"server/pkg/validation"
)

func AuthRoutes(
	api fiber.Router,
	db *gorm.DB,
	v *validation.CustomValidator,
	jwt utils.JwtService,
) {
	v.RegisterValidation("password", validation.PasswordValidation)

	repository := NewRepository(db)
	emailSender := email.NewEmailSender(config.AppConfig)
	service := NewService(repository, jwt, emailSender)
	handler := NewHandler(service)

	router := api.Group("/auth")

	router.Use(func(c fiber.Ctx) error {
		c.Locals("validator", v)
		return c.Next()
	})

	// Public Routes
	router.Post("/register", handler.RegisterUser)
	router.Post("/login", handler.LoginUser)
	router.Post("/refresh-token", handler.RefreshToken)
	router.Post("/forgot-password", handler.ForgotPassword)
	router.Post("/verify-otp", handler.VerifyOtp)
	router.Post("/reset-password", handler.ResetPassword)
	router.Post("/verify-email", handler.VerifyEmail)
	router.Post("/send-verification", handler.SendVerificationOtp) // also supported as public if email in body

	// Authenticated Routes
	router.Get("/me",
		middlewares.Authentication(jwt),
		handler.GetMe,
	)
	router.Post("/logout",
		middlewares.Authentication(jwt),
		handler.LogoutUser,
	)
	router.Post("/change-password",
		middlewares.Authentication(jwt),
		handler.ChangePassword,
	)
}
