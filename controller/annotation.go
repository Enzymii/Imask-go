package controller

import (
	"github.com/labstack/echo/v4"
	"imask-go/model"
	"net/http"
)

type AnnotationBody struct {
	Json   string `json:"json"`
	TaskId uint64 `json:"taskId"`
}

type SetAnnoStatBody struct {
	ID     uint64 `json:"id"`
	Status uint64 `json:"stat"`
}

func CreateAnnotation(c echo.Context) error {
	username := Auth(c)
	if username == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"msg": "Unauthorized"})
	}

	var user model.User
	res := model.DB.First(&user, "username = ?", username)
	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": "Internal Server Error"})
	}

	body := &AnnotationBody{}
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"msg": "Bad Request"})
	}

	res = model.DB.Create(&model.Annotation{
		Json:     body.Json,
		TaskID:   body.TaskId,
		AuthorId: username,
		Author:   user,
	})
	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": "Internal Server Error"})
	}

	return c.JSON(http.StatusOK, map[string]string{"msg": "OK"})
}

func GetAnnotation(c echo.Context) error {
	username := Auth(c)
	if username == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"msg": "Unauthorized"})
	}

	task := c.QueryParam("task")
	author := c.QueryParam("author")

	if task == "" && author == "" {
		author = username
	}

	if task != "" {
		var annotations []model.Annotation
		res := model.DB.Where("task_id = ?", task).Find(&annotations)
		if res.Error != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"msg": "Internal Server Error"})
		}

		return c.JSON(http.StatusOK, annotations)
	} else {
		var annotations []model.Annotation
		res := model.DB.Where("author_id = ?", author).Find(&annotations)
		if res.Error != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"msg": "Internal Server Error"})
		}

		return c.JSON(http.StatusOK, annotations)
	}
}

func UpdateAnnotationStatus(c echo.Context) error {
	username := Auth(c)
	if username == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"msg": "Unauthorized"})
	}

	param := &SetAnnoStatBody{}
	if err := c.Bind(param); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res := model.DB.Model(&model.Annotation{}).Where("id = ?", param.ID).Update("status", param.Status)
	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": "Internal Server Error"})
	}

	return c.JSON(http.StatusOK, map[string]string{"msg": "OK"})
}
