package handler

import (
	"bookstore/app/database"
	"bookstore/app/driver"
	middle "bookstore/app/middleware"
	"bookstore/http/handler/routes"

	"github.com/labstack/echo/v4"
)

func SetUpRoutes(e *echo.Echo) {
	connection, _ := driver.ConnectMysql()

	// Import seeder
	bookSeeder := database.BookMigration(connection)
	bookSeeder.ImportSeeder()

	userSeeder := database.UserMigration(connection)
	userSeeder.ImportSeeder()

	// Routes for bookstore
	boostore := routes.BooksHandler(connection)
	e.GET("/books", boostore.GetAll, middle.Auth, middle.Roles("books:view"))
	e.GET("/books/:id", boostore.GetById, middle.Auth, middle.Roles("books:view"))
	e.POST("/books", boostore.Create, middle.Auth, middle.Roles("books:create"))
	e.PUT("/books", boostore.Update, middle.Auth, middle.Roles("books:update"))
	e.DELETE("/books/:id", boostore.Delete, middle.Auth, middle.Roles("books:delete"))

	// Routes for users
	users := routes.UsersHandler(connection)
	e.GET("/users", users.GetAll)
}
