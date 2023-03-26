package models

import "gorm.io/gorm"

type Ratting struct {
	gorm.Model
	Ratting  float32 `json:"ratting"`
	MoovieID uint    `json:"moovie_id"`
}

type ResponseRatting struct {
	ID       uint
	Ratting  float32
	MoovieID uint
}
