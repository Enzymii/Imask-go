package controller

import (
	"github.com/labstack/echo/v4"
	"imask-go/model"
	"imask-go/utils"
	"net/http"
	"time"
)

func checkPassword(username string, password string) bool {
	var user model.User

	res := model.DB.First(&user, "username = ?", username)
	if res.Error != nil {
		return false
	}

	return utils.CheckPassword(password, user.Password)
}

type LoginModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c echo.Context) error {
	loginMod := new(LoginModel)
	if err := c.Bind(loginMod); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	res := checkPassword(loginMod.Username, loginMod.Password)

	if !res {
		err := c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid username or password",
		})
		return err
	}

	c.SetCookie(&http.Cookie{
		Name:     "username",
		Value:    loginMod.Username,
		Domain:   "127.0.0.1",
		Expires:  time.Now().Add(time.Second * utils.COOKIE_EXPIRE_TIME),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
	err := c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login success",
	})
	return err
}
