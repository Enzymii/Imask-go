package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetStatus(c echo.Context) error {
	username := Auth(c)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "ok",
		"username": username,
	})
}
