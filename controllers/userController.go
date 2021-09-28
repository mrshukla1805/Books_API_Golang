package controllers

import (
	"BOOKS_API/database"
	"BOOKS_API/middlewares"
	"BOOKS_API/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AllAuthors(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "authors"); err != nil {
		return err
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(database.DB, &models.Author{}, page))
}

func CreateAuthor(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "authors"); err != nil {
		return err
	}

	var author models.Author

	if err := c.BodyParser(&author); err != nil {
		return err
	}

	author.SetPassword("1234")

	database.DB.Create(&author)

	return c.JSON(author)
}

func GetAuthor(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	author := models.Author{
		Id: uint(id),
	}

	database.DB.Preload("Role").Find(&author)

	return c.JSON(author)
}

func UpdateAuthor(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	author := models.Author{
		Id: uint(id),
	}

	if err := c.BodyParser(&author); err != nil {
		return err
	}

	database.DB.Model(&author).Updates(author)

	return c.JSON(author)
}

func DeleteAuthor(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	author := models.Author{
		Id: uint(id),
	}

	database.DB.Delete(&author)

	return nil
}
