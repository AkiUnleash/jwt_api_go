package controllers

import (
	"jwt/domain/database"
	"jwt/domain/models"
	"jwt/middle"
	"net/http"

	"github.com/labstack/echo/v4"
)

func DiaryWrite(c echo.Context) error {

	// パラメータのBodyからデータをBind
	var data map[string]string
	if err := c.Bind(&data); err != nil {
		return err
	}

	// CookieからUIDを取得
	uid, err := middle.CurrentUserUid(c)
	if err != nil {
		return err
	}

	// DBにデータ登録
	diary := models.Diary{
		Uid:  uid,
		Body: data["body"],
	}
	database.DB.Create(&diary)

	return c.JSON(http.StatusOK, diary)
}
