package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/lennyochanda/LiveOak/tokenutil"
)

func JWTMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")

		if len(t) == 2 {
			authToken := t[1]
			if _, err := tokenutil.IsValid(authToken, secret); err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": err.Error(),
				})
				c.Abort()
			}; 

			userId, err := tokenutil.ExtractID(authToken, secret)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": err.Error(),
				})
				c.Abort()
			}
			
			fmt.Print("id", userId)
			c.Set("liveoak-id", userId)
		}
		c.Next()
	}
}
