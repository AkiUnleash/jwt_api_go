package jwt

import (
	"jwt/config"
	"time"

	"jwt/domain/models"

	"github.com/labstack/echo/v4"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(c echo.Context, account models.Account) (string, error) {
	// ヘッダーとペイロードを設定
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    account.Uid,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := claims.SignedString([]byte(config.Config.Secretkey))

	return token, err
}

func ReadToken(c echo.Context, cookie string) (string, error) {

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Config.Secretkey), nil
		})

	if err != nil {
		return "error", err
	}

	// トークンからクレームを取得
	claims := token.Claims.(*jwt.StandardClaims)
	err = nil
	return claims.Issuer, err
}
