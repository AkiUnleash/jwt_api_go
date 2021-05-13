package cookie

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

const CookieName = "jwt"

func SetCookie(c echo.Context, token string) {
	// Cookieに保存
	cookie := new(http.Cookie)
	cookie.Domain = "http://localhost:5000"
	cookie.Path = "/"
	cookie.Name = CookieName
	cookie.Value = token
	cookie.Expires = time.Now().Add(time.Hour * 24)
	cookie.HttpOnly = true
	cookie.Secure = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
}

func DeleteCookie(c echo.Context) error {
	// Cookieに保存
	cookie, err := c.Cookie(CookieName)
	if err != nil {
		return err
	}
	cookie.Domain = "http://localhost:5000"
	cookie.Path = "/"
	cookie.MaxAge = -1
	cookie.HttpOnly = true
	cookie.Secure = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)

	return nil
}
