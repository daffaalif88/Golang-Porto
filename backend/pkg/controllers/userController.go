package controllers

import (
	"golang-porto/backend/pkg/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// CreateUser creates a new user
func CreateUser(c echo.Context) error {
	var user models.User
	// Bind JSON payload ke model User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// Simpan data user ke database
	if err := db.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}

	// Kirimkan respons sukses
	return c.JSON(http.StatusCreated, user)
}

// GetUser retrieves all users
func GetUsers(c echo.Context) error {
	var users []models.User
	// Ambil semua data user beserta data terkait
	if err := db.Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve users"})
	}
	return c.JSON(http.StatusOK, users)
}

// GetUserByID retrieves a user by ID
func GetUserByID(c echo.Context) error {
	var user models.User
	id := c.Param("id")
	// Temukan user berdasarkan ID beserta data terkait
	if err := db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve user"})
	}
	return c.JSON(http.StatusOK, user)
}

// UpdateUser updates a user by ID
func UpdateUser(c echo.Context) error {
	var user models.User
	id := c.Param("id")
	// Temukan user berdasarkan ID
	if err := db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve user"})
	}

	// Bind JSON payload ke model User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// Simpan perubahan ke database
	if err := db.Save(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update user"})
	}
	return c.JSON(http.StatusOK, user)
}

// DeleteUser deletes a user by ID
func DeleteUser(c echo.Context) error {
	var user models.User
	id := c.Param("id")
	// Temukan user berdasarkan ID
	if err := db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve user"})
	}

	// Hapus user dari database
	if err := db.Delete(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete user"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "User deleted"})
}
