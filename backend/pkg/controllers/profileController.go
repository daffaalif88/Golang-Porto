package controllers

import (
	"golang-porto/backend/pkg/config"
	"golang-porto/backend/pkg/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// createProfile tambah data
func CreateProfile(c echo.Context) error {
	var profile models.Profile
	if err := c.Bind(&profile); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input "})
	}

	if err := db.Create(&profile).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create profile"})
	}

	return c.JSON(http.StatusOK, profile)
}

func GetProfiles(c echo.Context) error {
	var profiles []models.Profile
	if err := config.DB.Find(&profiles).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve profiles"})
	}
	return c.JSON(http.StatusOK, profiles)
}

func GetProfileByID(c echo.Context) error {
	var profile models.Profile
	id := c.Param("id")
	if err := db.First(&profile, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "profile not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve profile"})
	}
	return c.JSON(http.StatusOK, profile)
}

func UpdateProfile(c echo.Context) error {
	var profile models.Profile
	id := c.Param("id")
	if err := db.First(&profile, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "profile not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve profile"})
	}

	// Bind JSON payload ke model profile
	if err := c.Bind(&profile); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// Simpan perubahan ke database
	if err := db.Save(&profile).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update profile"})
	}
	return c.JSON(http.StatusOK, profile)
}

func DeleteProfile(c echo.Context) error {
	var profile models.Profile
	id := c.Param("id")
	if err := db.First(&profile, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "profile not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve profile"})
	}

	// Hapus profile dari database
	if err := db.Delete(&profile).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete profile"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "profile deleted"})
}
