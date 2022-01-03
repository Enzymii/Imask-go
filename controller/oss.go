package controller

import (
	"crypto/hmac"
	"crypto/sha1"
	b64 "encoding/base64"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/labstack/echo/v4"
	"imask-go/model"
	"imask-go/utils"
	"io"
	"net/http"
	"time"
)

type File struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func UploadFinished(c echo.Context) error {
	username := Auth(c)
	if username == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"msg": "Unauthorized"})
	}

	var user model.User
	res := model.DB.First(&user, "username = ?", username)
	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": "Internal Server Error"})
	}
	param := &[]File{}
	if err := c.Bind(param); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	for _, file := range *param {
		model.DB.Create(&model.Media{
			Name:    file.Name,
			Type:    file.Type,
			Owner:   user,
			OwnerID: user.Username,
		})
	}
	return c.JSON(http.StatusOK, map[string]string{"msg": "success"})
}

func GetCollection(c echo.Context) error {
	username := Auth(c)
	if username == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"msg": "Unauthorized"})
	}
	var media []model.Media
	res := model.DB.Find(&media, "owner_id = ?", username)
	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": "Internal Server Error"})
	}
	var mediaList []File
	for i := range media {
		mediaList = append(mediaList, File{Name: media[i].Name, Type: media[i].Type})
	}

	if len(mediaList) == 0 {
		return c.JSON(http.StatusOK, make([]string, 0))
	}
	return c.JSON(http.StatusOK, mediaList)
}

func GetUploadSignature(c echo.Context) error {
	if Auth(c) == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"msg": "Unauthorized"})
	}

	expireTime := time.Now().Add(10 * time.Minute).Format("2006-01-02T15:04:05.000Z")
	policy := `{"expiration":"` + expireTime +
		`","conditions":[{"bucket":"imask-media"},["starts-with", "$key", "media/"]]}`
	policyB64 := b64.StdEncoding.EncodeToString([]byte(policy))
	keyForSign := utils.OSS_ACCESS_KEY_SECRET
	h := hmac.New(sha1.New, []byte(keyForSign))
	_, err := io.WriteString(h, policyB64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	Signature := b64.StdEncoding.EncodeToString(h.Sum(nil))

	return c.JSON(http.StatusOK, map[string]string{
		"OSSAccessKeyId": utils.OSS_ACCESS_KEY_ID,
		"policy":         policyB64,
		"Signature":      Signature,
	})
}

func GetDownloadURL(c echo.Context) error {
	if Auth(c) == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"msg": "Unauthorized"})
	}

	filename := c.QueryParam("filename")

	client, err := oss.New(utils.OSS_ENDPOINT, utils.OSS_ACCESS_KEY_ID, utils.OSS_ACCESS_KEY_SECRET)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	bucket, err := client.Bucket(utils.OSS_BUCKET_NAME)
	objectName := "media/" + filename
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	signedURL, err := bucket.SignURL(objectName, oss.HTTPGet, utils.OSS_EXPIRE_TIME)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"url": signedURL,
	})
}
