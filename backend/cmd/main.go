package main

import (
	"golang-porto/backend/pkg/config"
	"golang-porto/backend/pkg/controllers"
	"golang-porto/backend/pkg/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Buat instance Echo
	e := echo.New()

	// Inisialisasi database
	db := config.ConnectDatabase()

	// Inisialisasi controller dengan database
	controllers.Initialize(db)

	// Middleware CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	// Protected routes here
	routes.SetupRoutes(e)

	// Serve static files
	e.Static("/static", "pkg/static")

	// Run the server
	e.Start(":8080")
}
