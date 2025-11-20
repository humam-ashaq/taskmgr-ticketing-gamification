package controllers

import (
	"strings"
	"ticketing-backend/database"
	"ticketing-backend/models"

	"github.com/gofiber/fiber/v2"
)

func CreateProject(c *fiber.Ctx) error {

	userID := c.Locals("user_id").(uint)

	var input struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Input tidak valid"})
	}

	slug := strings.ReplaceAll(strings.ToLower(input.Name), " ", "-")

	project := models.Project{
		Name:        input.Name,
		Slug:        slug,
		Description: input.Description,
		CreatedBy:   userID,
	}

	if err := database.DB.Create(&project).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal membuat project. Nama mungkin duplikat."})
	}

	member := models.ProjectMember{
		ProjectID: project.ID,
		UserID:    userID,
		Role:      "project_manager",
	}
	database.DB.Create(&member)

	return c.JSON(fiber.Map{
		"message": "Project berhasil dibuat!",
		"project": project,
	})
}

func GetMyProjects(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	var projects []models.Project

	database.DB.Joins("JOIN project_members on project_members.project_id = projects.id").
		Where("project_members.user_id = ?", userID).
		Find(&projects)

	return c.JSON(projects)
}
