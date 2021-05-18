package model

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Title string `gorm:"type:varchar(100);not null"`
	Body string `gorm:"type:longtext;not null"`
	Describe string `gorm:"type:varchar(255);not null"`
	WriteTime string `gorm:"type:varchar(100);not null"`
	ReadingTime string `gorm:"type:varchar(100);not null"`
	BackgroundImage string `gorm:"type:varchar(100);not null"`
}