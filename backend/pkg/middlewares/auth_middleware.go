package middlewares

import (
	"golang-porto/backend/pkg/utils"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// AuthenticationMiddleware checks if the user has a valid JWT token in Echo
func AuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get the token from the "Authorization" header
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Missing authentication token"})
		}

		// The token should be prefixed with "Bearer "
		tokenParts := strings.Split(tokenString, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid authentication token"})
		}

		tokenString = tokenParts[1]

		// Verify the token
		claims, err := utils.VerifyToken(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid authentication token"})
		}

		// Store the user_id in the context so it can be used in the request
		c.Set("user_id", claims["user_id"])

		// Proceed with the next handler
		return next(c)
	}
}
