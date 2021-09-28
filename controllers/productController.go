package controllers

import (
	"BOOKS_API/database"
	"BOOKS_API/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AllBooks(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(database.DB, &models.Book{}, page))
}

func CreateBook(c *fiber.Ctx) error {
	var book models.Book

	if err := c.BodyParser(&book); err != nil {
		return err
	}

	database.DB.Create(&book)

	return c.JSON(book)
}

func GetBook(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	book := models.Book{
		Id: uint(id),
	}

	database.DB.Find(&book)

	return c.JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	book := models.Book{
		Id: uint(id),
	}

	if err := c.BodyParser(&book); err != nil {
		return err
	}

	database.DB.Model(&book).Updates(book)

	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	book := models.Book{
		Id: uint(id),
	}

	database.DB.Delete(&book)

	return nil
}
