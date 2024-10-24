package handlers

import (
	"golang-porto/backend/pkg/models"
	"golang-porto/backend/pkg/utils"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

var db *gorm.DB

// Initialize mengatur koneksi database untuk controller
func Initialize(dbConn *gorm.DB) {
	db = dbConn
}

// Function for logging in
func Login(c echo.Context) error {
	var user models.User

	// Check user credentials and generate a JWT token
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid data"})
	}

	// Check if credentials are valid (replace this logic with real authentication)
	if user.Email == "user" && user.Password == "password" {
		// Generate a JWT token
		token, err := utils.GenerateToken(user.ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error generating token"})
		}

		return c.JSON(http.StatusOK, map[string]string{"token": token})
	}

	return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
}

// Function for registering a new user (for demonstration purposes)
func Register(c echo.Context) error {
	var user models.User

	// Bind JSON body ke struct user
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid data"})
	}

	// Simpan data user ke database
	if err := db.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}

	// Buat profil yang terkait dengan user yang baru dibuat
	profile := models.Profile{
		UserID: user.ID, // pastikan field ini ada di struct Profile
		// tambahkan field lain sesuai kebutuhan, seperti default avatar atau bio kosong
	}

	// Simpan profil ke database
	if err := db.Create(&profile).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create profile"})
	}

	// Kirim response sukses
	return c.JSON(http.StatusCreated, map[string]string{"message": "User and profile registered successfully"})
}
