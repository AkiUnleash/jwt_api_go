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
