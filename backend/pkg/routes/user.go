package routes

import (
	"golang-porto/backend/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func SetupUserRoutes(e *echo.Group) {
	e.GET("/users", controllers.GetUsers)
	e.POST("/users", controllers.CreateUser)
	e.GET("/users/:id", controllers.GetUserByID)
	e.PUT("/users/:id", controllers.UpdateUser)
	e.DELETE("/users/:id", controllers.DeleteUser)
}
