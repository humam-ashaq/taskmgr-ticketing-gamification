package controllers

import (
	"os"
	"ticketing-backend/database"
	"ticketing-backend/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Input tidak valid"})
	}

	if data["password"] == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Password wajib diisi"})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Name:         data["name"],
		Username:     data["username"],
		Email:        data["email"],
		PasswordHash: string(password),
		Role:         "developer",
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal membuat user. Email/Username mungkin sudah dipakai."})
	}

	stats := models.UserStats{
		UserID: user.ID,
		XP:     0,
		Level:  1,
	}
	database.DB.Create(&stats)

	return c.JSON(fiber.Map{
		"message": "Sukses mendaftar!",
		"user":    user,
	})
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Input tidak valid"})
	}

	var user models.User

	database.DB.Preload("Stats").Where("email = ?", data["email"]).First(&user)

	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "User tidak ditemukan"})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(data["password"]))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Password salah!"})
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "rahasia_dapur_capstone"
	}

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal login"})
	}

	return c.JSON(fiber.Map{
		"message": "Login Berhasil!",
		"token":   t,
		"user":    user,
	})
}
