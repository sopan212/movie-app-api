package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/movie-app-api/config"
	"github.com/movie-app-api/models"
	"gorm.io/gorm/clause"
)

type ReqCateMoo struct {
	CategoryID  uint
	MoovieID    uint
	Description string
}

func CategoryByMoovie(c *gin.Context) {
	var reqCateMoo ReqCateMoo

	err := c.ShouldBindJSON(&reqCateMoo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"meesage": "bad request",
			"error":   err.Error(),
		})
		return
	}

	reqCategory := models.CategoryMoovie{
		CategoryID: reqCateMoo.CategoryID,
		MoovieID:   reqCateMoo.MoovieID,
		Decription: reqCateMoo.Description,
	}

	insert := config.DB.Create(&reqCategory)
	if insert.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
			"error":   insert.Error.Error(),
		})

		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"data":    reqCategory,
	})
}

func GetMoovieByCategoryID(c *gin.Context) {
	id := c.Param("id")
	var category []models.Categorie

	err := c.BindJSON(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
			"error":   err.Error(),
		})
		return
	}

	config.DB.Preload(clause.Associations).First(&category, "id = ?", id)
	c.JSON(http.StatusOK, gin.H{
		"message": "get data success",
		"data":    category,
	})
}

func GetCategoryMoovie(c *gin.Context) {
	categoryMoovie := []models.CategoryMoovie{}

	config.DB.Preload(clause.Associations).Find(&categoryMoovie)

	responsecategorymoovie := []models.ResponseCategoryMoovie{}
	for _, respcm := range categoryMoovie {
		resp := models.ResponseCategoryMoovie{
			Id:           respcm.ID,
			CategoryName: respcm.Categorie.CategoryName,
			MoovieName:   respcm.Moovie.Title,
			Description:  respcm.Decription,
			CreatedAt:    respcm.CreatedAt,
		}
		responsecategorymoovie = append(responsecategorymoovie, resp)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    responsecategorymoovie,
	})
}
