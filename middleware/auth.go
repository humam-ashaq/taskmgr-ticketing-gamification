package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Protected(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{"error": "Akses ditolak, token tidak ada"})
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "rahasia_dapur_capstone"
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return c.Status(401).JSON(fiber.Map{"error": "Token tidak valid"})
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(float64)

	c.Locals("user_id", uint(userID))

	return c.Next()
}
