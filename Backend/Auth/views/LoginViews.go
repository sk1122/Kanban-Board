package views

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"Backend/Auth/helper"
)

// Login User Data for Binding
type LoginUserData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginUser(c *gin.Context) {
	var u LoginUserData

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	// Check If the User is Registered
	isUserAuthenticated := helper.CompareUsers(u.Username, u.Password)

	if isUserAuthenticated {

		access_token, refresh_token, err := helper.CreateToken(u.Username)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errosr": err.Error()})
			return
		}

		tokens := map[string]string{
			"access_token": access_token,
			"refresh_token": refresh_token,
		}

		c.JSON(http.StatusOK, gin.H{"token": tokens})
		return
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Login Info Is Wrong"})
	}
}