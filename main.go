package main

import (
	"weddingproject/database"
	"weddingproject/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize database
	database.InitDatabase()

	// Create Echo instance
	e := echo.New()

	// Routes for clients
	e.GET("/clients", handlers.GetClients)
	e.GET("/clients/:id", handlers.GetClientByID)
	e.POST("/clients", handlers.CreateClient)
	e.DELETE("/clients/:id", handlers.DeleteClient)

	// Routes for events
	e.GET("/events", handlers.GetEvents)
	e.GET("/events/:id", handlers.GetEventByID)
	e.POST("/events", handlers.CreateEvent)
	e.DELETE("/events/:id", handlers.DeleteEvent)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
