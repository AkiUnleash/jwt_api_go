package database

import (
	"fmt"
	"jwt/config"
	"jwt/domain/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Connet() {

	DBMS := config.Config.Dbms
	USER := config.Config.User
	PASS := config.Config.Pass
	PROTOCOL := config.Config.Protocol
	DBNAME := config.Config.Dbname

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME

	connection, err := gorm.Open(DBMS, CONNECT)

	fmt.Println(CONNECT)

	if err != nil {
		panic("could not connet to the database")
	}

	DB = connection

	// テーブル自動生成
	connection.AutoMigrate(&models.Account{})
	connection.AutoMigrate(&models.Diary{})
}
