package routes

import (
	"golang-porto/backend/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func SetupProjectRoutes(e *echo.Group) {
	e.GET("/projects", controllers.GetProjects)
	e.POST("/projects", controllers.CreateProject)
	e.GET("/projects/:id", controllers.GetProjectByID)
	e.PUT("/projects/:id", controllers.UpdateProject)
	e.DELETE("/projects/:id", controllers.DeleteProject)
}
