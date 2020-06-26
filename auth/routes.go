package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Register the main input to for a example route
func Register(group string, router *echo.Echo) {
	router.GET("/", Registration)
	router.POST("/", Registration)
}

// Registration Auth for default authentication
func Registration(c *echo.Context) {
	// log.Print("Registered default route.")
	c.String(http.StatusOK, "Auth API")
}
