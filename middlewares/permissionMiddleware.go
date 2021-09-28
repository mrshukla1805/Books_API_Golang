package middlewares

import (
	"errors"
	"fmt"
	"strconv"

	"BOOKS_API/database"

	"BOOKS_API/models"

	"BOOKS_API/util"

	"github.com/gofiber/fiber/v2"
)

func IsAuthorized(c *fiber.Ctx, page string) error {
	cookie := c.Cookies("jwt")

	Id, err := util.ParseJwt(cookie)

	fmt.Println(Id)

	if err != nil {
		return err
	}

	authorId, _ := strconv.Atoi(Id)

	author := models.Author{
		Id: uint(authorId),
	}

	database.DB.Preload("Role").Find(&author)

	role := models.Role{
		Id: author.RoleId,
	}

	database.DB.Preload("Permissions").Find(&role)

	fmt.Println(role.Permissions)
	if c.Method() == "GET" {
		for _, permission := range role.Permissions {
			if permission.Name == "view_"+page || permission.Name == "edit_"+page {
				return nil
			}
		}
	} else {
		for _, permission := range role.Permissions {
			fmt.Println(permission.Name)
			if permission.Name == "edit_"+page {
				return nil
			}
		}
	}

	c.Status(fiber.StatusUnauthorized)
	return errors.New("unauthorized")
}
