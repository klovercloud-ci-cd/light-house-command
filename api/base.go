package api

import (
	v1 "github.com/klovercloud/lighthouse-command/api/v1"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Routes(e *echo.Echo) {
	// Index Page
	e.GET("/", index)

	// Health Page
	e.GET("/health", health)
	v1.Router(e.Group("/api/v1"))
}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "This is KloverCloud light house command service")
}

func health(c echo.Context) error {
	return c.String(http.StatusOK, "I am live!")
}
