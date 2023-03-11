package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/movie-app-api/config"
	"github.com/movie-app-api/models"
	"gorm.io/gorm/clause"
)

func GetMoovieByCategories(c *gin.Context) {

	var categories []models.Categorie

	// config.DB.Find(&moovies)
	config.DB.Preload(clause.Associations).Find(&categories)

	c.JSON(200, gin.H{
		"message": "success",
		"data":    categories,
	})
}
func GetMoovieById(c *gin.Context) {
	id := c.Param("id")
	var moovies models.Moovie

	data := config.DB.First(&moovies, id)
	if data.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
			"status":  "failed get data",
			"error":   data.Error.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    moovies,
	})
}
