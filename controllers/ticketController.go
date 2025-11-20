package controllers

import (
	"ticketing-backend/database"
	"ticketing-backend/models"

	"github.com/gofiber/fiber/v2"
)

func CreateTicket(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var input struct {
		ProjectID     uint   `json:"project_id"`
		Title         string `json:"title"`
		Description   string `json:"description"`
		Priority      string `json:"priority"`
		EstimateHours int    `json:"estimate_hours"`
		AssigneeID    *uint  `json:"assignee_id"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Input tidak valid"})
	}

	if input.ProjectID == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Project ID wajib diisi"})
	}

	ticket := models.Ticket{
		ProjectID:     input.ProjectID,
		Title:         input.Title,
		Description:   input.Description,
		Priority:      input.Priority,
		EstimateHours: input.EstimateHours,
		CreatorID:     userID,
		AssigneeID:    input.AssigneeID,
		Status:        "todo",
	}

	if err := database.DB.Create(&ticket).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal membuat ticket"})
	}

	return c.JSON(fiber.Map{
		"message": "Ticket berhasil dibuat!",
		"ticket":  ticket,
	})
}

func GetTickets(c *fiber.Ctx) error {
	projectID := c.Query("project_id")

	var tickets []models.Ticket

	query := database.DB.Preload("Creator").Preload("Assignee")

	if projectID != "" {
		query = query.Where("project_id = ?", projectID)
	}

	query.Find(&tickets)

	return c.JSON(tickets)
}
