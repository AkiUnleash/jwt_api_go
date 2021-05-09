package controllers

import (
	"jwt/domain/database"
	"jwt/domain/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Account(c echo.Context) error {

	// CookieのデータからUserを取得
	var account models.Account
	database.DB.Where("id = ?", 2).First(&account)

	// JSONを返す
	return c.JSON(http.StatusOK, account)
}
