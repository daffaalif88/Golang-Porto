package main

import (
	"golang-porto/backend/pkg/config"
	"golang-porto/backend/pkg/controllers"
	"golang-porto/backend/pkg/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// Buat instance Echo
	e := echo.New()

	// Inisialisasi database
	db := config.ConnectDatabase()

	// Inisialisasi controller dengan database
	controllers.Initialize(db)

	// Protected routes (require authentication)

	// Protected routes here
	routes.SetupRoutes(e)

	// Serve static files
	e.Static("/static", "pkg/static")

	// Run the server
	e.Start(":8080")
}
