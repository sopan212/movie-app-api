package main

import (
	"net/http"
	"github.com/movie-app-api/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.initDB()
	route := gin.Default()
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	route.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
