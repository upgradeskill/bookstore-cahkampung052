package handler

import (
	"bookstore/app/handler/routes"

	"github.com/labstack/echo/v4"
)

func SetUpRoutes(e *echo.Echo) {
	// Routes for bookstore
	boostore := routes.BooksHandler{}
	e.GET("/books", boostore.GetAll)
	e.GET("/books/:id", boostore.GetById)
	e.POST("/books", boostore.Create)
	e.PUT("/books", boostore.Update)
	e.DELETE("/books", boostore.Delete)

	// Routes for users
	users := routes.UsersHandler{}
	e.GET("/users", users.GetAll)
}
