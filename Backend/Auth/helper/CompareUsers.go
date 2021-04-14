package helper

import (
	"Backend/Auth/db"
)

func CompareUsers(username string, password string) bool {
	var u db.User

	if err := db.DB.Where("Username = ?", username).First(&u).Error; err != nil {
		return false;
	}

	return CheckPasswordHash(password, u.Password)
}