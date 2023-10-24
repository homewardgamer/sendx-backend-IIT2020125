package api

import (
	"github.com/gin-gonic/gin"
)

func IsPayingCustomer() gin.HandlerFunc {
	return func(c *gin.Context) {
		isPaying := c.Query("isPaying") // Assuming "isPaying" is a query parameter

		if isPaying == "true" {
			c.Set("isPaying", true)
		} else {
			c.Set("isPaying", false)
		}

		c.Next()
	}
}
