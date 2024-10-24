package routes

import (
	"golang-porto/backend/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func SetupProfileRoutes(e *echo.Group) {
	e.GET("/profiles", controllers.GetProfiles)
	e.POST("/profiles", controllers.CreateProfile)
	e.GET("/profiles/:id", controllers.GetProfileByID)
	e.PUT("/profiles/:id", controllers.UpdateProfile)
	e.DELETE("/profiles/:id", controllers.DeleteProfile)
}
