package views

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"Backend/Auth/helper"
)

func Demo(c *gin.Context) {
	_, err := helper.TokenValid(c.Request)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token Invalid"})
	}

	c.JSON(http.StatusOK, gin.H{"data": "Authorized"})
}