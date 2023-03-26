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
	var moovies []models.Moovie

	data := config.DB.Preload(clause.Associations).Find(&moovies, id)
	if data.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
			"status":  "failed get data",
			"error":   data.Error.Error(),
		})
		c.Abort()
		return
	}

	respMoovieID := []models.ResponseMoovies{}
	for _, respM := range moovies {
		respCategor := []models.ResponseCategory{}
		for _, respC := range respM.Categories {
			resp := models.ResponseCategory{
				ID:           respC.CategoryID,
				Categoryname: respC.Categorie.CategoryName,
			}
			respCategor = append(respCategor, resp)
		}
		respReviews := []models.ResponseReview{}
		respRatting := []models.ResponseRatting{}
		for _, rattR := range respM.Rating {
			resp := models.ResponseRatting{
				ID:       rattR.ID,
				Ratting:  rattR.Ratting,
				MoovieID: rattR.MoovieID,
			}
			respRatting = append(respRatting, resp)
		}
		for _, respR := range respM.Reviews {
			respRe := models.ResponseReview{
				ID:     respR.ID,
				Review: respR.Review,
			}
			respReviews = append(respReviews, respRe)
		}
		resp := models.ResponseMoovies{
			ID:          respM.ID,
			Title:       respM.Title,
			Description: respM.Description,
			Years:       respM.Years,
			Poster:      respM.Poster,
			Rating:      respRatting,
			Categories:  respCategor,
			Reviews:     respReviews,
		}
		respMoovieID = append(respMoovieID, resp)
	}
	c.JSON(200, gin.H{
		"message": "success",
		"data":    respMoovieID,
	})
}
