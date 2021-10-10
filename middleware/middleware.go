package middleware

import (
	"fmt"
	"net/http"
	"strings"

	helper "northwind-api/helpers"

	"github.com/gin-gonic/gin"
)

// Authz validates token and authorizes users
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization header provided")})
			c.Abort()
			return
		}

		token := strings.Split(clientToken, "Bearer ")[1]

		claims, err := helper.ValidateToken(token)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_Name)
		c.Set("user_id", claims.User_Id)
		c.Set("last_name", claims.Last_Name)
		c.Set("uid", claims.UID)

		c.Next()

	}
}
