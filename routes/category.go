package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/movie-app-api/config"
	"github.com/movie-app-api/models"
	"gorm.io/gorm/clause"
)

func GetCategory(c *gin.Context) {
	var categories []models.Categorie

	config.DB.Preload(clause.Associations).Find(&categories)

	c.JSON(http.StatusOK, gin.H{
		"message": "get success",
		"status":  "success",
		"data":    categories,
	})
}

func GetCategoryByID(c *gin.Context) {
	id := c.Param("id")
	var categorie models.Categorie

	config.DB.First(&categorie, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "get success",
		"status":  "success",
		"data":    categorie,
	})
}
