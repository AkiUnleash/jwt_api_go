package main

import (
	"jwt/domain/database"
	"jwt/domain/routing"
)

// @title JWT-login-example
// @version 1.0.0
// @description API for login processing using JWT. Developed in Go langage.

// @license.name MIT
// @license.url https://github.com/tcnksm/tool/blob/master/LICENCE

// @host http://localhost:8081
// @BasePath /
func main() {

	database.Connet()

	routing.Routing()

}
