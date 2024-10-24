package routes

import (
	"golang-porto/backend/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func SetupPortofolioHomeRoutes(e *echo.Echo) {
	// Rute untuk halaman HTML
	e.GET("/portofolio", controllers.ShowPortofolio)
}
