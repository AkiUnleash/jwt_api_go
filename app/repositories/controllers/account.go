package controllers

import (
	"jwt/domain/database"
	"jwt/domain/models"
	"jwt/until/cookie"
	"jwt/until/jwt"
	"jwt/until/uid"
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
		Uid:      uid.Generate(),
		Email:    data["email"],
		Password: password,
	}

	database.DB.Create(&account)

	return c.JSON(http.StatusOK, account)
}

// ログイン処理
func Login(c echo.Context) error {

	// パラメータからデータを抽出
	var data map[string]string
	if err := c.Bind(&data); err != nil {
		return err
	}

	var account models.Account

	// アカウント登録済みを確認する
	database.DB.Where("email = ?", data["email"]).First(&account)
	if account.Id == 0 {
		return c.JSON(http.StatusNotFound, "404 Not Found")
	}

	// パスワード不一致
	if err := bcrypt.CompareHashAndPassword(account.Password, []byte(data["password"])); err != nil {
		return c.JSON(http.StatusBadRequest, "incrrect password")
	}

	// JSON Web Tokenの作成
	token, err := jwt.GenerateToken(c, account)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "unauthenticated")
	}

	// Cookie保存
	cookie.SetCookie(c, token)

	// 結果出力
	return c.JSON(http.StatusOK, "success")
}
