package middlewares

import (
	httpresponse "server/pkg/httpResponse"
	"server/pkg/utils"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

func Authentication(jwtService utils.JwtService) fiber.Handler {
	return func(c fiber.Ctx) error {
		// Get token from cookie (check access_token and token) or Authorization header
		tokenString := c.Cookies("access_token")
		if tokenString == "" {
			tokenString = c.Cookies("token")
		}
		if tokenString == "" {
			authHeader := c.Get("Authorization")
			if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
				tokenString = authHeader[7:]
			}
		}

		if tokenString == "" {
			return c.Status(http.StatusUnauthorized).JSON(httpresponse.Error{
				Success: false,
				Message: "Authentication required",
			})
		}

		// Validate token
		claims, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(httpresponse.Error{
				Success: false,
				Message: "Authentication failed. Invalid or expired token",
			})
		}

		// Store user information in context
		c.Locals("user_id", claims.UserID)
		c.Locals("name", claims.Name)
		c.Locals("email", claims.Email)
		c.Locals("role", claims.Role)

		return c.Next()
	}
}
