package controllers

import (
	"golang-porto/backend/pkg/config"
	"golang-porto/backend/pkg/models"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

// RenderIndexPage menampilkan halaman index.html
func RenderIndexPage(c echo.Context) error {
	var users []models.User
	// Ambil semua data user beserta data terkait
	if err := config.DB.
		Preload("Projects").
		Preload("Skills").
		Preload("Contacts").
		Preload("Educations").
		Preload("Experiences").
		Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve users"})
	}

	// Siapkan data untuk template
	data := map[string]interface{}{
		"Users": users,
	}

	// Render template HTML
	tmpl, err := template.ParseFiles("pkg/views/home/index.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to parse template"})
	}
	return tmpl.Execute(c.Response().Writer, data)
}