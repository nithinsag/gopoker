package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var db *gorm.DB

func InitDB() *gorm.DB {
	var err error

	db, err = gorm.Open("sqlite3", viper.GetString("dbname"))
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println(db)
	MigrateUsers()
	return db

}

func CloseDB() {
	db.Close()
}
