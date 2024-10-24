package controllers

import (
	"golang-porto/backend/pkg/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// CreateEducation creates a new education
func CreateEducation(c echo.Context) error {
	var education models.Education
	// Bind JSON payload ke model Education
	if err := c.Bind(&education); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// Simpan data education ke database
	if err := db.Create(&education).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create education"})
	}

	// Kirimkan respons sukses
	return c.JSON(http.StatusCreated, education)
}

// GetEducations retrieves all educations
func GetEducations(c echo.Context) error {
	var educations []models.Education
	// Ambil semua data education
	if err := db.Find(&educations).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve educations"})
	}
	return c.JSON(http.StatusOK, educations)
}

// GetEducationByID retrieves an education by ID
func GetEducationByID(c echo.Context) error {
	var education models.Education
	id := c.Param("id")
	// Temukan education berdasarkan ID
	if err := db.First(&education, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Education not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve education"})
	}
	return c.JSON(http.StatusOK, education)
}

// UpdateEducation updates an education by ID
func UpdateEducation(c echo.Context) error {
	var education models.Education
	id := c.Param("id")
	// Temukan education berdasarkan ID
	if err := db.First(&education, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Education not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve education"})
	}

	// Bind JSON payload ke model Education
	if err := c.Bind(&education); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// Simpan perubahan ke database
	if err := db.Save(&education).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update education"})
	}
	return c.JSON(http.StatusOK, education)
}

// DeleteEducation deletes an education by ID
func DeleteEducation(c echo.Context) error {
	var education models.Education
	id := c.Param("id")
	// Temukan education berdasarkan ID
	if err := db.First(&education, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Education not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve education"})
	}

	// Hapus education dari database
	if err := db.Delete(&education).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete education"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Education deleted"})
}
