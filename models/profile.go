package models

import (
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	Age    int    `json:"age"`
	Gender string `json:"gender"`
	UserID uint   `json:"user_id"`
}
