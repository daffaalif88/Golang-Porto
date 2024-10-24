package routes

import (
	"golang-porto/backend/pkg/middlewares"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	// Public route (no middleware)
	SetupHandlerRoutes(e) // This one remains public without authentication
	SetupPortofolioHomeRoutes(e)
	// Grouping protected routes
	protected := e.Group("")                            // Create a group for protected routes
	protected.Use(middlewares.AuthenticationMiddleware) // Apply authentication middleware to the group

	// Protected routes
	SetupUserRoutes(protected)
	SetupProjectRoutes(protected)
	SetupSkillRoutes(protected)
	SetupContactRoutes(protected)
	SetupEducationRoutes(protected)
	SetupExperienceRoutes(protected)

	SetupProfileRoutes(protected)
}
