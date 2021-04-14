package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"Backend/Auth/helper"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := helper.TokenValid(c.Request)

		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()

			return
		}

		c.Next()
	}
}