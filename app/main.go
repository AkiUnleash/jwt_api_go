package main

import (
	"jwt/domain/database"
	"jwt/domain/routing"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	database.Connet()

	routing.Routing()

}
