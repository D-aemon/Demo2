package model

import "github.com/jinzhu/gorm"

type Relation struct {
	gorm.Model
	ArticleId uint
	TagId uint
}
