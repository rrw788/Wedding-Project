package handlers

import (
	"net/http"
	"time"
	"weddingproject/database"
	"weddingproject/models"
	"weddingproject/utils"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// JWT Secret Key
var jwtSecret = []byte("your_secret_key")

// LoginHandler authenticates the user and generates a JWT.
func LoginHandler(c echo.Context) error {
	// Parse request body
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	// Find user in database
	var user models.User
	if err := database.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"})
	}

	// Check password
	if !utils.CheckPassword(user.Password, input.Password) {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"})
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
}
