package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	JanCode string
	Name    string
	Price   int
}
