package controller

import (
	"github.com/labstack/echo/v4"
	"imask-go/model"
	"net/http"
)

type CreateReqBody struct {
	Name  string `json:"name"`
	Files string `json:"files"`
}

func CreateTask(c echo.Context) error {
	username := Auth(c)
	if username == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"msg": "Unauthorized"})
	}

	var user model.User
	res := model.DB.First(&user, "username = ?", username)
	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": "Internal Server Error"})
	}
	param := &CreateReqBody{}
	if err := c.Bind(param); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res = model.DB.Create(&model.Task{
		Name:     param.Name,
		AuthorId: username,
		Author:   user,
		Content:  param.Files,
	})

	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": "Internal Server Error"})
	}

	return c.JSON(http.StatusOK, map[string]string{"msg": "OK"})
}

func GetTasks(c echo.Context) error {
	username := Auth(c)
	if username == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"msg": "Unauthorized"})
	}

	var tasks []model.Task
	res := model.DB.Find(&tasks)
	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": "Internal Server Error"})
	}

	return c.JSON(http.StatusOK, tasks)
}

func GetMyTasks(c echo.Context) error {
	username := Auth(c)
	if username == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"msg": "Unauthorized"})
	}

	var tasks []model.Task
	res := model.DB.Find(&tasks, "author_id = ?", username)
	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": "Internal Server Error"})
	}

	return c.JSON(http.StatusOK, tasks)
}
