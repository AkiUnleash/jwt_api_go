package routing

import (
	"jwt/repositories/controllers"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "jwt/docs"
)

func Routing() {
	e := echo.New()

	e.GET("account/user", controllers.Account)
	e.POST("account/signup", controllers.Register)
	e.POST("account/login", controllers.Login)
	e.POST("account/logout", controllers.Logout)
	e.GET("account/nowuser", controllers.CurrentUser)

	e.GET("diary", controllers.DiaryRead)
	e.POST("diary", controllers.DiaryWrite)
	e.DELETE("diary/:id", controllers.DiaryDelete)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8081"))
}
