package main

import (
    "log"
    "github.com/gofiber/fiber/v2"
    "github.com/LewisCai/SimpleCRMProject/database"
    "github.com/LewisCai/SimpleCRMProject/lead"
)

func main() {
    // Initialize the database
    err := database.Connect()
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    defer func() {
        err := database.Disconnect()
        if err != nil {
            log.Fatal("Failed to disconnect from database:", err)
        }
    }()

    app := fiber.New()

    // Serve static files from the "static" directory
    app.Static("/", "./static")

    // Set up API routes
    app.Get("/api/v1/lead", lead.GetLeads)
    app.Post("/api/v1/lead", lead.NewLead)
    app.Delete("/api/v1/lead/:id", lead.DeleteLead)

    // Start the server
    log.Fatal(app.Listen(":3000"))
}
