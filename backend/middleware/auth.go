package middleware

import (
	"fmt"
	"os"
	"strings"

	"askwise.com/m/v2/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserEmailKey contextKey = "userEmail"

func AuthRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(utils.Error("Missing Authorization header"))
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(utils.Error("Invalid Authorization header"))
		}

		tokenString := parts[1]
		secret := os.Getenv("NEXTAUTH_SECRET")
		if secret == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(utils.Error("Missing NEXTAUTH_SECRET"))
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(utils.Error("Invalid or expired token"))
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(utils.Error("Invalid claims"))
		}

		email, ok := claims["email"].(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(utils.Error("Email not found in token"))
		}

		c.Locals(string(UserEmailKey), email)

		return c.Next()
	}
}

func GetUserEmail(c *fiber.Ctx) string {
	return c.Locals(string(UserEmailKey)).(string)
}
