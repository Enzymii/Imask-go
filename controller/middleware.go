package controller

import (
	"github.com/labstack/echo/v4"
	"imask-go/model"
)

func Auth(c echo.Context) string {
	cookie, err := c.Cookie("username")
	if err != nil {
		return ""
	}
	if cookie.Name != "username" {
		return ""
	}
	cookieName := cookie.Value
	res := model.DB.First(&model.User{}, "username = ?", cookieName)
	if res.Error != nil || res.RowsAffected == 0 {
		return ""
	}
	return cookieName
}
