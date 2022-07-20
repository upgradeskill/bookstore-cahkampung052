package handler

import (
	controller "bookstore/http/handler/controller"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetUpRoutes(e *echo.Echo, connection *gorm.DB) {
	// Routes for bookstore
	boostore := controller.BooksHandler(connection)
	e.GET("/books", boostore.GetAll, JwtMiddleware, Auth("books:view"))
	e.GET("/books/:id", boostore.GetById, JwtMiddleware, Auth("books:view"))
	e.POST("/books", boostore.Create, JwtMiddleware, Auth("books:create"))
	e.PUT("/books", boostore.Update, JwtMiddleware, Auth("books:update"))
	e.DELETE("/books/:id", boostore.Delete, JwtMiddleware, Auth("books:delete"))

	// Routes for users
	users := controller.UsersHandler(connection)
	e.GET("/users", users.GetAll)
}
