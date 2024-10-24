package controllers

import (
	"golang-porto/backend/pkg/config"
	"golang-porto/backend/pkg/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Show data Portofolio
func ShowPortofolio(c echo.Context) error {
	var profiles []models.Profile
	if err := config.DB.
		Preload("Projects").
		Preload("Skills").
		Preload("Contacts").
		Preload("Educations").
		Preload("Experiences").
		Find(&profiles).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve profiles"})
	}
	return c.JSON(http.StatusOK, profiles)
}
