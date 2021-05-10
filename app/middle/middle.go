package middle

import (
	"jwt/until/cookie"
	"jwt/until/jwt"

	"github.com/labstack/echo/v4"
)

// ログイン中のUIDを取得
func CurrentUserUid(c echo.Context) (string, error) {
	// Cookieを取得
	cookie, err := c.Cookie(cookie.CookieName)
	if err != nil {
		return "", err
	}

	// Tokenを解析
	uid, err := jwt.ReadToken(c, cookie.Value)
	if err != nil {
		return "", err
	}

	return uid, err
}
