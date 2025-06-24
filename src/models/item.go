package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	JanCode string
	Name    string
	Price   int
}

func (Item *Item) Create(janCode string, name string, price int) {
	Item.JanCode = janCode
	Item.Name = name
	Item.Price = price
}

func (Item *Item) Read(janCode string) {
	Item.JanCode = janCode
}
