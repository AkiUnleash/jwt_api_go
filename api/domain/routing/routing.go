package routing

import (
	"jwt/repositories/controllers"

	_ "github.com/go-sql-driver/mysql"

	_ "jwt/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Routing() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"http://localhost:5000"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXCSRFToken},
	}))

	e.GET("account/user", controllers.Account)
	e.POST("account/signup", controllers.Register)
	e.POST("account/login", controllers.Login)
	e.POST("account/logout", controllers.Logout)
	e.GET("account/nowuser", controllers.CurrentUser)

	e.GET("diary", controllers.DiaryRead)
	e.POST("diary", controllers.DiaryWrite)
	e.DELETE("diary/:id", controllers.DiaryDelete)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8082"))
}
