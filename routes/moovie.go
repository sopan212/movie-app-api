package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/movie-app-api/config"
	"github.com/movie-app-api/models"
	"gorm.io/gorm/clause"
)

func GetMoovie(c *gin.Context) {
	var moovies []models.Moovie

	config.DB.Preload(clause.Associations).Find(&moovies)
	// config.DB.Preload(clause.Associations).Find(&moovies)

	respMoovie := []models.ResponseMoovies{}
	for _, Moov := range moovies {
		respCategory := []models.ResponseCategory{}
		for _, Ctgry := range Moov.Categories {
			respCategry := models.ResponseCategory{
				ID:           Ctgry.ID,
				Categoryname: Ctgry.Categorie.CategoryName,
			}
			respCategory = append(respCategory, respCategry)
		}
		respRatting := []models.ResponseRatting{}
		for _, rattR := range Moov.Rating {
			resp := models.ResponseRatting{
				ID:       rattR.ID,
				Ratting:  rattR.Ratting,
				MoovieID: rattR.MoovieID,
			}
			respRatting = append(respRatting, resp)
		}
		RespReview := []models.ResponseReview{}
		for _, review := range Moov.Reviews {
			respView := models.ResponseReview{
				ID:     review.ID,
				Review: review.Review,
			}
			RespReview = append(RespReview, respView)
		}
		moovieResp := models.ResponseMoovies{
			ID:          Moov.ID,
			Title:       Moov.Title,
			Description: Moov.Description,
			Years:       Moov.Years,
			Poster:      Moov.Poster,
			Rating:      respRatting,
			Categories:  respCategory,
			Reviews:     RespReview,
		}
		respMoovie = append(respMoovie, moovieResp)
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    respMoovie,
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

	c.JSON(http.StatusCreated, gin.H{
		"message":     "insert success",
		"ID":          moovie.ID,
		"Title":       moovie.Title,
		"Description": moovie.Description,
		"Years":       moovie.Years,
		"Poster":      moovie.Poster,
		"Rating":      moovie.Rating,
		"Category_ID": moovie.CategorieID,
	})
}

// // type reqMoovie struct {
// // 	Title       string `json:"title"`
// // 	Description string `json:"description"`
// // 	Years       int    `json:"years"`
// // 	Poster      string `json:"poster"`
// // 	CategoryID  uint   `json:"category_id"`
// }

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
		"message":     "update success",
		"id":          moovie.ID,
		"title":       moovie.Title,
		"description": moovie.Description,
		"years":       moovie.Years,
		"poster":      moovie.Poster,
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
