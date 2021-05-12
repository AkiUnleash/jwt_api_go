package controllers

import (
	"jwt/domain/database"
	"jwt/domain/models"
	"jwt/middle"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Login
// @Summary Diary registratinon process.
// @Description Can be executed only at login.
// @tags diary
// @Accept  json
// @Produce  json
// @Success 200
// @Router /diary [post]
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

// Login
// @Summary Processing to display diary
// @Description Can be executed only at login.
// @tags diary
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Diary
// @Router /diary [Get]
func DiaryRead(c echo.Context) error {

	// CookieからUIDを取得
	uid, err := middle.CurrentUserUid(c)
	if err != nil {
		return err
	}

	var diary []models.Diary
	database.DB.Where("uid = ?", uid).Order("created_at DESC").Find(&diary)

	return c.JSON(http.StatusOK, diary)
}

// Login
// @Summary Process to delete diary.
// @Description Can be executed only at login.
// @tags diary
// @Accept  json
// @Produce  json
// @Success 200
// @Router /diary/:id [Delete]
func DiaryDelete(c echo.Context) error {
	id := c.Param("id")
	var diary []models.Diary
	database.DB.Where("id = ?", id).Delete(&diary)

	return c.JSON(http.StatusOK, "OK")
}
