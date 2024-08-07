package main

import (
	"github.com/LewisCai/SimpleCRMProject/database"
	"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App) {
	app.Get(GetLeads)
	app.Get(GetLead)
	app.Post(NewLead)
	app.Delete(DeleteLead)
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