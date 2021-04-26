package model

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Title string `gorm:"type:varchar(100);not null"`
	Tag string `gorm:"type:varchar(100);not null"`
	Body string `gorm:"type:varchar(255);not null"`
}
