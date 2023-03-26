package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/movie-app-api/auth"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "UnAuthorization",
				"error":   http.StatusUnauthorized,
			})
			c.Abort()
			return
		}
		_, err := auth.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
				"status":  http.StatusUnauthorized,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
				"status":  http.StatusUnauthorized,
			})
			c.Abort()
			return
		}
		role, err := auth.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
				"status":  http.StatusUnauthorized,
			})
			c.Abort()
			return
		}

		//role 1 = admin
		if role != 1 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
				"status":  http.StatusUnauthorized,
			})
			c.Abort()
			return
		}

		//c.Set("x-email", email)
		c.Next()
	}
}
