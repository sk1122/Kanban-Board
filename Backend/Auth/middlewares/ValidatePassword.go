package middlewares

import (
	"net/http"
	"github.com/gin-gonic/gin"
    "unicode"

    "Backend/Auth/db"
)

func ValidatePassword() gin.HandlerFunc {
	return func(c *gin.Context) {

		var (
	        hasMinLen  = false
	        hasLower   = false
	        hasNumber  = false
	        hasSpecial = false
	    )
	    
	    var user db.User

	    if err := c.ShouldBindJSON(&user); err == nil {

	    	s := user.Password

		    if len(s) >= 7 {
		        hasMinLen = true
		    }
		    for _, char := range s {
		        switch {
		        case unicode.IsLower(char):
		            hasLower = true
		        case unicode.IsNumber(char):
		            hasNumber = true
		        case unicode.IsPunct(char) || unicode.IsSymbol(char):
		            hasSpecial = true
		        }
		    }
		    if !hasMinLen || !hasLower || !hasNumber || !hasSpecial {
		    	c.JSON(http.StatusBadRequest, gin.H{"error": "Password Doesn't meet criteria"})
		    	c.Abort()

		    	return
		    }
		
		}

		c.Next()
	}
}