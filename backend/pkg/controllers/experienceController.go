package controllers

import (
	"golang-porto/backend/pkg/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// CreateExperience creates a new experience
func CreateExperience(c echo.Context) error {
	var experience models.Experience
	// Bind JSON payload ke model Experience
	if err := c.Bind(&experience); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// Simpan data experience ke database
	if err := db.Create(&experience).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create experience"})
	}

	// Kirimkan respons sukses
	return c.JSON(http.StatusCreated, experience)
}

// GetExperiences retrieves all experiences
func GetExperiences(c echo.Context) error {
	var experiences []models.Experience
	// Ambil semua data experience
	if err := db.Find(&experiences).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve experiences"})
	}
	return c.JSON(http.StatusOK, experiences)
}

// GetExperienceByID retrieves an experience by ID
func GetExperienceByID(c echo.Context) error {
	var experience models.Experience
	id := c.Param("id")
	// Temukan experience berdasarkan ID
	if err := db.First(&experience, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Experience not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve experience"})
	}
	return c.JSON(http.StatusOK, experience)
}

// UpdateExperience updates an experience by ID
func UpdateExperience(c echo.Context) error {
	var experience models.Experience
	id := c.Param("id")
	// Temukan experience berdasarkan ID
	if err := db.First(&experience, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Experience not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve experience"})
	}

	// Bind JSON payload ke model Experience
	if err := c.Bind(&experience); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// Simpan perubahan ke database
	if err := db.Save(&experience).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update experience"})
	}
	return c.JSON(http.StatusOK, experience)
}

// DeleteExperience deletes an experience by ID
func DeleteExperience(c echo.Context) error {
	var experience models.Experience
	id := c.Param("id")
	// Temukan experience berdasarkan ID
	if err := db.First(&experience, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Experience not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve experience"})
	}

	// Hapus experience dari database
	if err := db.Delete(&experience).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete experience"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Experience deleted"})
}
