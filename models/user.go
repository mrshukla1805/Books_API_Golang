package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Author struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	Password  []byte `json:"-"`
	RoleId    uint   `json:"role_id"`
	Role      Role   `json:"role" gorm:"foreignKey:RoleId"`
}

func (author *Author) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	author.Password = hashedPassword
}

func (author *Author) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(author.Password, []byte(password))
}

func (author *Author) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Author{}).Count(&total)

	return total
}

func (author *Author) Take(db *gorm.DB, limit int, offset int) interface{} {
	var books []Author

	db.Preload("Role").Offset(offset).Limit(limit).Find(&books)

	return books
}
