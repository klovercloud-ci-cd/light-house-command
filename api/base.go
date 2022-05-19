package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Routes(e *echo.Echo) {
	// Index Page
	e.GET("/", index)

	// Health Page
	e.GET("/health", health)
}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "This is KloverCloud light house command service")
}

func health(c echo.Context) error {
	return c.String(http.StatusOK, "I am live!")
}
