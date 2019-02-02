package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// Db connection
var Db *gorm.DB

// InitDB Initializes database
func InitDB() *gorm.DB {
	var err error

	Db, err = gorm.Open("sqlite3", viper.GetString("dbname"))
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println(Db)

	return Db

}

// CloseDB closes Db connection
func CloseDB() {
	Db.Close()
}
