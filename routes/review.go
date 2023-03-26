package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/movie-app-api/config"
	"github.com/movie-app-api/models"
	"gorm.io/gorm/clause"
)

func AddReview(c *gin.Context) {

	var review models.Review

	err := c.BindJSON(&review)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
			"err":     err.Error(),
		})
		c.Abort()
		return
	}

	data := config.DB.Create(&review)

	if data.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "insert failed",
			"err":     data.Error.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "insert success",
		"data":    review,
	})
}

func GetReviews(c *gin.Context) {
	var reviews []models.Review

	data := config.DB.Preload(clause.Associations).Find(&reviews)
	if data.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed get data",
			"err":     data.Error.Error(),
		})

		c.Abort()
		return
	}

	Respreviews := []models.ResponseReview{}
	for _, review := range reviews {
		resp := models.ResponseReview{
			ID:     review.ID,
			Review: review.Review,
		}
		Respreviews = append(Respreviews, resp)
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "get review succes",
		"data":    Respreviews,
	})
}
