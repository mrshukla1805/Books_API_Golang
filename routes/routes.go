package routes

import (
	"BOOKS_API/controllers"

	"BOOKS_API/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use(middlewares.IsAuthenticated)

	app.Put("/api/users/info", controllers.UpdateInfo)
	app.Put("/api/users/password", controllers.UpdatePassword)

	app.Get("/api/user", controllers.Author)
	app.Post("/api/logout", controllers.Logout)

	app.Get("/api/users", controllers.AllAuthors)
	app.Post("/api/users", controllers.CreateAuthor)
	app.Get("/api/users/:id", controllers.GetAuthor)
	app.Put("/api/users/:id", controllers.UpdateAuthor)
	app.Delete("/api/users/:id", controllers.DeleteAuthor)

	app.Get("/api/roles", controllers.AllRoles)
	app.Post("/api/roles", controllers.CreateRole)
	app.Get("/api/roles/:id", controllers.GetRole)
	app.Put("/api/roles/:id", controllers.UpdateRole)
	app.Delete("/api/roles/:id", controllers.DeleteRole)

	app.Get("/api/permissions", controllers.AllPermissions)

	app.Get("/api/products", controllers.AllBooks)
	app.Post("/api/products", controllers.CreateBook)
	app.Get("/api/products/:id", controllers.GetBook)
	app.Put("/api/products/:id", controllers.UpdateBook)
	app.Delete("/api/products/:id", controllers.DeleteBook)

	app.Post("/api/upload", controllers.Upload)
	app.Static("/api/uploads", "./uploads")

	app.Get("/api/orders", controllers.AllOrders)

}
