package controllers

import (
	"jwt/config"
	"jwt/domain/database"
	"jwt/domain/models"
	"jwt/until/uid"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
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

	// JSON Web Token
	// ヘッダーとペイロードを設定
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    account.Uid,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := claims.SignedString([]byte(config.Config.Secretkey))

	if err != nil {
		return c.JSON(http.StatusUnauthorized, "unauthenticated")
	}

	// Cookieに保存
	cookie := new(http.Cookie)
	cookie.Name = "jwt"
	cookie.Value = token
	cookie.Expires = time.Now().Add(time.Hour * 24)
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, token)
}
