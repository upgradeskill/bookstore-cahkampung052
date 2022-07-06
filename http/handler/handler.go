package handler

import (
	"bookstore/app/database"
	"bookstore/app/driver"
	"bookstore/http/handler/routes"

	"github.com/labstack/echo/v4"
)

func SetUpRoutes(e *echo.Echo) {
	connection, _ := driver.ConnectMysql()

	// Import seeder
	migration := database.BookMigration(connection)
	migration.ImportSeeder()

	// Routes for bookstore
	boostore := routes.BooksHandler(connection)
	e.GET("/books", boostore.GetAll)
	e.GET("/books/:id", boostore.GetById)
	e.POST("/books", boostore.Create)
	e.PUT("/books", boostore.Update)
	e.DELETE("/books/:id", boostore.Delete)

	// Routes for users
	users := routes.UsersHandler{}
	e.GET("/users", users.GetAll)
}
