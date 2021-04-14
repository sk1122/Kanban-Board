package views

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"Backend/Auth/db"
	"Backend/Auth/helper"
)


// User Struct for Input
type UserData struct {
	Username string `json:"username"`
	Email 	 string `json:"email"`
	Password string `json:"password"`
}

func RegisterUser(c *gin.Context) {
	var u UserData

	// Check if POST JSON is in right format 
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hashing Password
	hashedPassword, err := helper.HashPassword(u.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hashing Error"})
		return
	}

	// Create User Model
	user := db.User{Username: u.Username, Email: u.Email, Password: hashedPassword}
	
	// Check If Username exists
	isUsernameExists := helper.CheckIfUserNameExists(user.Username)

	if isUsernameExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username Already Exists"})
		return
	}

	// Check If Email Exists
	isEmailExists := helper.CheckIfEmailExists(user.Email)

	if isEmailExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email Already Exists"})
		return
	}

	// Save Model To Database
	db.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": hashedPassword})
}