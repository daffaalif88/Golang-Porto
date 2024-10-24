package routes

import (
	"golang-porto/backend/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func SetupSkillRoutes(e *echo.Group) {
	e.GET("/skills", controllers.GetSkills)
	e.POST("/skills", controllers.CreateSkill)
	e.GET("/skills/:id", controllers.GetSkillByID)
	e.PUT("/skills/:id", controllers.UpdateSkill)
	e.DELETE("/skills/:id", controllers.DeleteSkill)
}
