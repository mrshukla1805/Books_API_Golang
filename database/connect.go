package database

import (
	"BOOKS_API/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open("root:mysql1805@tcp(127.0.0.1:3306)/go_admin?charset=utf8&parseTime=True"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}

	DB = database

	database.AutoMigrate(&models.Author{}, &models.Role{}, &models.Permission{}, &models.Book{}, &models.Order{}, &models.OrderItem{})
}
