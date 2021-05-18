package model

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model
	TagName string `gorm:"type:varchar(100);not null"`
}
