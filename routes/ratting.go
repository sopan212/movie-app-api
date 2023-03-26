package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/movie-app-api/config"
	"github.com/movie-app-api/models"
	"gorm.io/gorm/clause"
)

func InsertRatting(c *gin.Context) {
	var rating models.Ratting

	c.BindJSON(&rating)

	data := config.DB.Create(&rating)
	if data.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
			"error":   data.Error.Error(),
		})
		c.Abort()
		return
	}
	respRatting := []models.ResponseRatting{}
	var rattings []models.Ratting
	for _, ratt := range rattings {
		resp := models.ResponseRatting{
			ID:      ratt.ID,
			Ratting: ratt.Ratting,
		}
		respRatting = append(respRatting, resp)
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "insert success",
		"data":    respRatting,
	})
}

func GetRatting(c *gin.Context) {
	var rattings []models.Ratting

	data := config.DB.Preload(clause.Associations).Find(&rattings)
	if data.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
			"error":   data.Error.Error(),
		})
		c.Abort()
		return
	}
	respRatting := []models.ResponseRatting{}
	for _, ratt := range rattings {
		resp := models.ResponseRatting{
			ID:       ratt.ID,
			Ratting:  ratt.Ratting,
			MoovieID: ratt.MoovieID,
		}
		respRatting = append(respRatting, resp)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    respRatting,
	})
}
