package handlers

import (
	"net/http"
	"weddingproject/database"
	"weddingproject/models"

	"github.com/labstack/echo/v4"
)

// Get all events
func GetEvents(c echo.Context) error {
	var events []models.Event
	database.DB.Find(&events)
	return c.JSON(http.StatusOK, events)
}

// Get a single event by ID
func GetEventByID(c echo.Context) error {
	id := c.Param("id")
	var event models.Event
	if result := database.DB.First(&event, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Event not found"})
	}
	return c.JSON(http.StatusOK, event)
}

// Create a new event
func CreateEvent(c echo.Context) error {
	event := new(models.Event)
	if err := c.Bind(event); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	database.DB.Create(&event)
	return c.JSON(http.StatusCreated, event)
}

// Delete an event by ID
func DeleteEvent(c echo.Context) error {
	id := c.Param("id")
	var event models.Event
	if result := database.DB.Delete(&event, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Event not found"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Event deleted successfully"})
}
