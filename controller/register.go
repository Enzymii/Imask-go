package controller

import (
	"github.com/labstack/echo/v4"
	"imask-go/model"
	"imask-go/utils"
	"net/mail"
)

type RegisterModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func Register(c echo.Context) error {
	println("register")

	regMod := new(RegisterModel)
	if err := c.Bind(regMod); err != nil {
		return err
	}

	println(regMod.Username)

	_, emailErr := mail.ParseAddress(regMod.Email)
	if len(regMod.Username) < 6 || len(regMod.Password) < 6 || emailErr != nil {
		err := c.JSON(400, map[string]interface{}{
			"message": "Invalid params",
		})
		return err
	}

	exists := model.DB.First(&model.User{}, "username = ? or email = ?", regMod.Username, regMod.Email)

	if exists.RowsAffected != 0 {
		err := c.JSON(400, map[string]interface{}{
			"message": "Username or email already exists",
		})
		return err
	}

	encrypt, err := utils.Encrypt(regMod.Password)
	res := model.DB.Create(&model.User{
		Username: regMod.Username,
		Password: encrypt,
		Email:    regMod.Email,
	})

	if res.Error != nil {
		err := c.JSON(400, map[string]interface{}{
			"message": "Register failed",
		})
		return err
	}

	err = c.JSON(200, map[string]interface{}{
		"message": "Register success",
	})
	return err
}
