package models

import "gorm.io/gorm"

type Categorie struct {
	gorm.Model
	CategoryName string `json:"categoryname"`
	Moovies      []Moovie
}
