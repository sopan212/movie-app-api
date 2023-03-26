package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/movie-app-api/config"
	"github.com/movie-app-api/models"
	"gorm.io/gorm/clause"
)

func InsertCategory(c *gin.Context) {
	var category models.Categorie

	err := c.BindJSON(&category)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	data := config.DB.Create(&category)
	if data.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
			"error":   data.Error.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "insert success",
		"data":    category,
	})
}

func GetCategory(c *gin.Context) {
	categories := []models.CategoryMoovie{}

	err := config.DB.Preload(clause.Associations).Find(&categories)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "get data failed",
			"error":   err.Error.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "get category succes",
		"data":    categories,
	})
}

// func GetCategory(c *gin.Context) {
// 	var categories []models.Categorie

// 	config.DB.Preload(clause.Associations).Find(&categories)

// 	respByCategories := []models.ResponseForCategory{}
// 	for _, respCat := range categories {
// 		respMoo := []models.ResponseMoovies{}
// 		var movies []models.Moovie = make([]models.Moovie, len(respMoo))

// 		for i, resp := range respMoo {
// 			movies[i] = resp.Moovie
// 		}
// 		fmt.Println(movies)
// 		myStruct := MyStruct{
// 			movies:movies,

// 		// for _, respM := range respCat.Moovies {
// 		// 	rsp := models.ResponseMoovies{
// 		// 		ID:          respM.ID,
// 		// 		Title:       respM.Moovie.Title,
// 		// 		Description: respM.Moovie.Description,
// 		// 		Years:       respM.Moovie.Years,
// 		// 		Poster:      respM.Moovie.Poster,
// 		// 	}
// 		// 	respMoo = append(respMoo, rsp)

// 		// }
// 		respcA := models.ResponseForCategory{
// 			ID:           respCat.ID,
// 			Categoryname: respCat.CategoryName,
// 			Moovies:      respMoo,
// 		}
// 		respByCategories = append(respByCategories, respcA)
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "get success",
// 		"status":  "success",
// 		"data":    respByCategories,
// 	})
// }

func GetCategoryByID(c *gin.Context) {
	id := c.Param("id")
	var categorie []models.Categorie

	config.DB.Preload(clause.Associations).Find(&categorie, "id = ?", id)

	c.JSON(http.StatusOK, gin.H{
		"message": "get success",
		"status":  "success",
		"data":    categorie,
	})
}

// func GetCategoryByQuery(c *gin.Context) {
// 	query := c.Query("category")
// 	var categorie models.Categorie

// 	config.DB.Find(&categorie, query)

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "get success",
// 		"status":  "success",
// 		"data":    categorie,
// 	})
// }
