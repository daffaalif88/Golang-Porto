package routes

import (
	"golang-porto/backend/pkg/handlers"

	"github.com/labstack/echo/v4"
)

func SetupHandlerRoutes(e *echo.Echo) {
	e.POST("/login", handlers.Login)
	e.POST("/register", handlers.Register)
	e.POST("/logout", handlers.Logout)
}
