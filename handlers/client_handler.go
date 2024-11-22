package handlers

import (
	"log"
	"net/http"
	"weddingproject/database"
	"weddingproject/models"

	"github.com/labstack/echo/v4"
)

func GetClients(c echo.Context) error {
	var clients []models.Client
	database.DB.Find(&clients)
	return c.JSON(http.StatusOK, clients)
}

func GetClientByID(c echo.Context) error {
	id := c.Param("id")
	var client models.Client
	if result := database.DB.First(&client, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Client not found"})
	}
	return c.JSON(http.StatusOK, client)
}

func CreateClient(c echo.Context) error {
	client := new(models.Client)

	// Bind JSON ke struct
	if err := c.Bind(client); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
	}

	// Simpan data ke database
	if err := database.DB.Create(&client).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create client"})
	} else {
		log.Println("Creating client:", client)
	}

	return c.JSON(http.StatusCreated, client)
}

func DeleteClient(c echo.Context) error {
	id := c.Param("id")
	var client models.Client

	// Cari data client
	if result := database.DB.First(&client, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Client not found"})
	}

	// Hapus client jika ditemukan
	if err := database.DB.Delete(&client).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to delete client"})
	} else {
		log.Println("Deleting client ID:", id)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Client deleted successfully"})
}
