package main

import (
	"github.com/gin-gonic/gin"
	"github.com/movie-app-api/config"
	"github.com/movie-app-api/routes"
)

func main() {
	config.InitDB()
	route := gin.Default()

	v1 := route.Group("api/v1")
	{
		v1.GET("ping", getHome)

		mooviesAdmin := v1.Group("/moovies/admin")
		{
			mooviesAdmin.GET("/", routes.GetMoovie)
			mooviesAdmin.POST("/", routes.InsertMoovie)
			mooviesAdmin.PUT("/:id", routes.EditMoovie)
			mooviesAdmin.DELETE("/:id", routes.DeleteMoovie)

		}
		mooviesUser := v1.Group("/moovies/user")
		{
			mooviesUser.GET("/", routes.GetMoovie)
			mooviesUser.GET("/:id", routes.GetMoovieById)
			mooviesUser.GET("/categories", routes.GetMoovieByCategories)

		}
		moovieCategories := v1.Group("/moovies/user/categories")
		{
			moovieCategories.GET("/", routes.GetCategory)
			moovieCategories.GET("/:id", routes.GetCategory)

		}

	}

	route.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ini adalah halaman home",
	})
}
