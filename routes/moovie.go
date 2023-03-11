package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/movie-app-api/config"
	"github.com/movie-app-api/models"
)

func GetMoovie(c *gin.Context) {
	var moovies []models.Moovie

	config.DB.Find(&moovies)
	// config.DB.Preload(clause.Associations).Find(&moovies)

	c.JSON(200, gin.H{
		"message": "success",
		"data":    moovies,
	})
}

func InsertMoovie(c *gin.Context) {
	var moovie models.Moovie

	err := c.BindJSON(&moovie)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		c.Abort()
		return
	}

	data := config.DB.Create(&moovie)
	if data.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "insert failed",
			"status":  "bad request",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "insert success",
		"data":    moovie,
	})
}

func EditMoovie(c *gin.Context) {
	id := c.Param("id")
	var moovie models.Moovie
	var reqMoovie models.Moovie

	c.BindJSON(&reqMoovie)

	data := config.DB.Model(&moovie).Where("id = ?", id).Updates(reqMoovie)
	if data.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "update failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "update success",
		"data":    moovie,
	})
}

func DeleteMoovie(c *gin.Context) {

	id := c.Param("id")
	var moovie models.Moovie

	config.DB.Delete(&moovie, id)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"status":  "bad request",
	// 		"message": "failed delete",
	// 	})

	// 	return
	// }
	c.JSON(http.StatusOK, gin.H{
		"message": "deleted succes",
	})
}
