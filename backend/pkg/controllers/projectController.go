package controllers

import (
	"golang-porto/backend/pkg/config"
	"golang-porto/backend/pkg/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// createProject tambah data
func CreateProject(c echo.Context) error {
	var project models.Project
	if err := c.Bind(&project); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input "})
	}

	if err := db.Create(&project).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create project"})
	}

	return c.JSON(http.StatusOK, project)
}

func GetProjects(c echo.Context) error {
	var projects []models.Project
	if err := config.DB.Find(&projects).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve projects"})
	}
	return c.JSON(http.StatusOK, projects)
}

func GetProjectByID(c echo.Context) error {
	var project models.Project
	id := c.Param("id")
	if err := db.First(&project, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "project not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve project"})
	}
	return c.JSON(http.StatusOK, project)
}

func UpdateProject(c echo.Context) error {
	var project models.Project
	id := c.Param("id")
	if err := db.First(&project, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "project not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve project"})
	}

	// Bind JSON payload ke model project
	if err := c.Bind(&project); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// Simpan perubahan ke database
	if err := db.Save(&project).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update project"})
	}
	return c.JSON(http.StatusOK, project)
}

func DeleteProject(c echo.Context) error {
	var project models.Project
	id := c.Param("id")
	if err := db.First(&project, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "project not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve project"})
	}

	// Hapus project dari database
	if err := db.Delete(&project).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete project"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "project deleted"})
}
