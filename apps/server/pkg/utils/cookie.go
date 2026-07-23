package utils

import (
	"time"

	"github.com/gofiber/fiber/v3"
)

// SetAuthCookies sets cookies for access_token ("token"), refresh_token, and session_token.
func SetAuthCookies(c fiber.Ctx, accessToken, refreshToken, sessionToken string) {
	accessTokenDuration := 24 * time.Hour       // 1 day
	refreshTokenDuration := 30 * 24 * time.Hour  // 30 days

	// Access Token (as "access_token" cookie for backward-compatibility with middlewares)
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Path:     "/",
		HTTPOnly: true,
		Secure:   false, // set to false for local dev HTTP, set true in production SSL
		SameSite: "Lax",
		MaxAge:   int(accessTokenDuration.Seconds()),
	})

	// Refresh Token
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/",
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Lax",
		MaxAge:   int(refreshTokenDuration.Seconds()),
	})

	// Session Token
	c.Cookie(&fiber.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Path:     "/",
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Lax",
		MaxAge:   int(refreshTokenDuration.Seconds()),
	})
}

// ClearAuthCookies clears all auth-related cookies.
func ClearAuthCookies(c fiber.Ctx) {
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    "",
		Path:     "/",
		HTTPOnly: true,
		MaxAge:   -1,
	})
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Path:     "/",
		HTTPOnly: true,
		MaxAge:   -1,
	})
	c.Cookie(&fiber.Cookie{
		Name:     "session_token",
		Value:    "",
		Path:     "/",
		HTTPOnly: true,
		MaxAge:   -1,
	})
}
