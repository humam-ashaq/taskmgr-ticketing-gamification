package routes

import (
	"ticketing-backend/controllers"
	"ticketing-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Ticketing Backend is Running! 🚀")
	})

	api := app.Group("/api")

	api.Post("/register", controllers.Register)
	api.Post("/login", controllers.Login)

	api.Use(middleware.Protected)

	api.Post("/projects", controllers.CreateProject)
	api.Get("/projects", controllers.GetMyProjects)
	api.Get("/projects/:id", controllers.GetProjectDetail)

	api.Get("/tickets", controllers.GetTickets)
	api.Post("/tickets", controllers.CreateTicket)
	api.Put("/tickets/:id", controllers.UpdateTicket)

}
