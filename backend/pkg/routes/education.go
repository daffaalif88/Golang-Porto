package routes

import (
	"golang-porto/backend/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func SetupEducationRoutes(e *echo.Group) {
	e.GET("/educations", controllers.GetEducations)
	e.POST("/educations", controllers.CreateEducation)
	e.GET("/educations/:id", controllers.GetEducationByID)
	e.PUT("/educations/:id", controllers.UpdateEducation)
	e.DELETE("/educations/:id", controllers.DeleteEducation)
}
