package helper 

import (
	"Backend/Auth/db"
)

func CheckIfUserNameExists(username string) bool {
	var u db.User

	if err := db.DB.Where("Username = ?", username).First(&u).Error; err != nil {
		return false
	} else {
		return true
	}
}

func CheckIfEmailExists(email string) bool {
	var u db.User

	if err := db.DB.Where("Email = ?", email).First(&u).Error; err != nil {
		return false
	} else {
		return true
	}
}