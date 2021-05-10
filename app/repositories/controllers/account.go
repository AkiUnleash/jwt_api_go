package controllers

import (
	"jwt/domain/database"
	"jwt/domain/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// アカウント表示
func Account(c echo.Context) error {

	// CookieのデータからUserを取得
	var account models.Account
	database.DB.Where("id = ?", 2).First(&account)

	// JSONを返す
	return c.JSON(http.StatusOK, account)
}

// アカウントの登録
func Register(c echo.Context) error {
	var data map[string]string

	if err := c.Bind(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	account := models.Account{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	database.DB.Create(&account)

	return c.JSON(http.StatusOK, account)
}
