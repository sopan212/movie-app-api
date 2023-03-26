package main

import (
	"github.com/gin-gonic/gin"
	"github.com/movie-app-api/config"
	"github.com/movie-app-api/routes"
)

func main() {
	config.InitDB()
	route := gin.Default()

	v1 := route.Group("api/v1/moovies")
	{
		v1.GET("ping", getHome)
		user := v1.Group("/user")
		{
			user.POST("/register", routes.RegisterUser)
			user.POST("/login", routes.GenerateToken)
			user.GET("/get", routes.GetUser)

		}
		userProfile := v1.Group("/user")
		{
			userProfile.POST("/profile", routes.CreateProfile)
			userProfile.GET("/profile", routes.GetProfile)

		}
		mooviesAdmin := v1.Group("/admin") //.Use(middlewares.IsAdmin())
		{
			mooviesAdmin.GET("/", routes.GetMoovie)
			mooviesAdmin.POST("/", routes.InsertMoovie)
			mooviesAdmin.PUT("/:id", routes.EditMoovie)
			mooviesAdmin.DELETE("/:id", routes.DeleteMoovie)
			mooviesAdmin.POST("/category", routes.InsertCategory)

		}
		mooviesUser := v1.Group("/user") //.Use(middlewares.Auth())
		{
			mooviesUser.GET("/", routes.GetMoovie)
			mooviesUser.GET("/:id", routes.GetMoovieById)
			mooviesUser.GET("/categories", routes.GetMoovieByCategories)

		}
		moovieCategories := v1.Group("/user/categories") //.Use(middlewares.Auth())
		{

			moovieCategories.GET("/", routes.GetCategory)
			moovieCategories.GET("/:id", routes.GetCategoryByID)
			//moovieCategories.GET("/category/:query", routes.GetCategoryByID)

		}
		moovieReview := v1.Group("/user/review") //.Use(middlewares.Auth())
		{

			moovieReview.POST("/", routes.AddReview)
			moovieReview.GET("/", routes.GetReviews)

		}
		moovieRatting := v1.Group("/user/ratting") //.Use(middlewares.Auth())
		{
			moovieRatting.POST("/", routes.InsertRatting)
			moovieRatting.GET("/", routes.GetRatting)
		}
		categoryMoovie := v1.Group("/user/categorymoovie") //.Use(middlewares.Auth())
		{
			categoryMoovie.POST("/", routes.CategoryByMoovie)
			categoryMoovie.GET("/", routes.GetCategoryMoovie)
			categoryMoovie.GET("/:id", routes.GetMoovieByCategoryID)
		}
	}

	route.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ini adalah halaman home",
	})
}
