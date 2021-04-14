package models 

import (
	"github.com/jinzhu/gorm"
)

type TodoData struct {
	gorm.Model
	Todo 	 string
	Category string
	Username string
}

type TodoDataList struct {
	gorm.Model
	Todo  string
	Doing string
	Done  string
	Trash string
	Username string
}