package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/LewisCai/SimpleCRMProject/database"
	"github.com/LewisCai/SimpleCRMProject/lead"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	// Connect to the database
	err := database.Connect()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
}

func main() {
	// Initialize the database
	initDatabase()
	defer func() {
		err := database.Disconnect()
		if err != nil {
			log.Fatal("Failed to disconnect from database:", err)
		}
	}()

	// Set up Fiber app and routes
	app := fiber.New()
	setupRoutes(app)

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
