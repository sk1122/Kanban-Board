package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "user.db")

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed To Connect Database")
		return
	}

	database.AutoMigrate(&User{})

	DB = database
}