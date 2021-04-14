package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB
func ConnectDatabase() {
	db, err := gorm.Open("sqlite3", "test.db")
    if err != nil {
        fmt.Println(err.Error())
        panic("failed to connect database")
    }

    // Migrate the schema
    db.AutoMigrate(&TodoData{})
    db.AutoMigrate(&TodoDataList{})
    DB = db
}