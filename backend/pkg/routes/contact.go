package routes

import (
	"golang-porto/backend/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func SetupContactRoutes(e *echo.Group) {
	e.GET("/contacts", controllers.GetContacts)
	e.POST("/contacts", controllers.CreateContact)
	e.GET("/contacts/:id", controllers.GetContactByID)
	e.PUT("/contacts/:id", controllers.UpdateContact)
	e.DELETE("/contacts/:id", controllers.DeleteContact)
}
