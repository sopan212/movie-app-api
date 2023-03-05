package config


import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
  )
  
	DB *gorm.io

  func initDB() {

	var err error
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/movie-app?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil{
		log.panic("failed to connect database")
	}
  }