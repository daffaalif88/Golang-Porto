package routes

import (
	"golang-porto/backend/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func SetupExperienceRoutes(e *echo.Group) {
	e.GET("/experiences", controllers.GetExperiences)
	e.POST("/experiences", controllers.CreateExperience)
	e.GET("/experiences/:id", controllers.GetExperienceByID)
	e.PUT("/experiences/:id", controllers.UpdateExperience)
	e.DELETE("/experiences/:id", controllers.DeleteExperience)
}
