package routing

import (
	"jwt/repositories/controllers"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
)

func Routing() {
	e := echo.New()

	e.GET("/user", controllers.Account)
	// e.POST("/login", controllers.Login)
	e.POST("/signup", controllers.Register)

	e.Logger.Fatal(e.Start(":8081"))
}
