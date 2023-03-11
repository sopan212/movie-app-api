package config

import (
	"github.com/movie-app-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	var err error
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/movie-app?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&models.Moovie{}, models.Categorie{})

	// DB.Create(&models.Categorie{
	// 	CategoryName: "drama",
	// })
	// DB.Create(&models.Categorie{
	// 	CategoryName: "comedy",
	// })
	// DB.Create(&models.Categorie{
	// 	CategoryName: "action",
	// })
	// DB.Create(&models.Categorie{
	// 	CategoryName: "drama",
	// })
	// DB.Create(&models.Categorie{
	// 	CategoryName: "anime",
	// })
}
