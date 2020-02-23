package db

import (
	"echo-gorm-mvc/config"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func Init() {
	configuration := config.GetConfig()
	connectString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true",
		configuration.DB_USERNAME, configuration.DB_PASSWORD, configuration.DB_HOST, configuration.DB_NAME)
	db, err = gorm.Open("mysql", connectString)

	// defer db.Close()
	if err != nil {
		panic(err)
	}
}

func DbManager() *gorm.DB {
	return db
}
