package routes

import (
	"golang-porto/backend/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func SetupHandlerRoutes(e *echo.Echo) {
	e.POST("/login", controllers.Login)
	e.POST("/register", controllers.Register)
}
