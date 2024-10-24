package controllers

import (
	"golang-porto/backend/pkg/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// CreateSkill creates a new skill
func CreateSkill(c echo.Context) error {
	var skill models.Skill
	// Bind JSON payload ke model Skill
	if err := c.Bind(&skill); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// Simpan data skill ke database
	if err := db.Create(&skill).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create skill"})
	}

	// Kirimkan respons sukses
	return c.JSON(http.StatusCreated, skill)
}

// GetSkills retrieves all skills
func GetSkills(c echo.Context) error {
	var skills []models.Skill
	// Ambil semua data skill
	if err := db.Find(&skills).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve skills"})
	}
	return c.JSON(http.StatusOK, skills)
}

// GetSkillByID retrieves a skill by ID
func GetSkillByID(c echo.Context) error {
	var skill models.Skill
	id := c.Param("id")
	// Temukan skill berdasarkan ID
	if err := db.First(&skill, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Skill not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve skill"})
	}
	return c.JSON(http.StatusOK, skill)
}

// UpdateSkill updates a skill by ID
func UpdateSkill(c echo.Context) error {
	var skill models.Skill
	id := c.Param("id")
	// Temukan skill berdasarkan ID
	if err := db.First(&skill, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Skill not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve skill"})
	}

	// Bind JSON payload ke model Skill
	if err := c.Bind(&skill); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// Simpan perubahan ke database
	if err := db.Save(&skill).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update skill"})
	}
	return c.JSON(http.StatusOK, skill)
}

// DeleteSkill deletes a skill by ID
func DeleteSkill(c echo.Context) error {
	var skill models.Skill
	id := c.Param("id")
	// Temukan skill berdasarkan ID
	if err := db.First(&skill, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Skill not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve skill"})
	}

	// Hapus skill dari database
	if err := db.Delete(&skill).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete skill"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Skill deleted"})
}
