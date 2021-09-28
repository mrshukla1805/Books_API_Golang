package models

import "gorm.io/gorm"

type Book struct {
	Id          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
}

func (book *Book) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Book{}).Count(&total)

	return total
}

func (book *Book) Take(db *gorm.DB, limit int, offset int) interface{} {
	var books []Book

	db.Offset(offset).Limit(limit).Find(&books)

	return books
}
