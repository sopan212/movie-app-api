package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/movie-app-api/config"
	"github.com/movie-app-api/models"
	"gorm.io/gorm/clause"
)

func CreateProfile(c *gin.Context) {
	var profile models.Profile

	c.BindJSON(&profile)
	config.DB.Create(&profile)

	c.JSON(http.StatusCreated, gin.H{"data": profile})
}

func GetProfile(c *gin.Context) {
	var user models.User

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
			"err":     err.Error(),
		})
		c.Abort()
		return
	}
	config.DB.Preload(clause.Associations).First(&user)

	c.JSON(http.StatusOK, gin.H{
		"id":     user.ID,
		"name":   user.Name,
		"gender": user.Profile.Gender,
		"age":    user.Profile.Age,
	})
}
