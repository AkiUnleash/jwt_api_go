package main

import (
	"jwt/domain/database"
	"jwt/domain/routing"
)

func main() {

	database.Connet()

	routing.Routing()

}
