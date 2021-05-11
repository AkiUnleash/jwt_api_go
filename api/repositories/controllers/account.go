package controllers

import (
	"jwt/domain/database"
	"jwt/domain/models"
	"jwt/middle"
	"jwt/until/cookie"
	"jwt/until/jwt"
	"jwt/until/uid"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Account(c echo.Context) error {

	// CookieのデータからUserを取得
	var account models.Account
	database.DB.Where("id = ?", 2).First(&account)

	// JSONを返す
	return c.JSON(http.StatusOK, account)
}

// CurentUser
// @Summary Show infomation about the currently logged in user.
// @Description Browse Account table.
// @tags account
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Account
// @Router /account/nowuser [get]
func CurrentUser(c echo.Context) error {

	// CookieからUIDを取得
	uid, err := middle.CurrentUserUid(c)
	if err != nil {
		return err
	}

	// uidからUserを取得
	var account models.Account
	database.DB.Where("uid = ?", uid).First(&account)

	// JSONを返す
	return c.JSON(http.StatusOK, account)

}

// Register
// @Summary Register Account infomation in the database.
// @Description Use the account table.
// @tags account
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Account
// @Router /account/signup [post]
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

// Login
// @Summary If the infomation passed in the request body matches the data in the table, a cookie will be issued.
// @Description JWT certification
// @tags account
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Account
// @Router /account/login [post]
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

// Login
// @Summary If the cookie exists, delete it.
// @Description JWT certification
// @tags account
// @Accept  json
// @Produce  json
// @Success 200
// @Router /account/logout [post]
func Logout(c echo.Context) error {

	// Cookie削除
	if err := cookie.DeleteCookie(c); err != nil {
		return c.JSON(http.StatusUnauthorized, "unauthenticated")
	}

	// 結果出力
	return c.JSON(http.StatusOK, "success")

}
