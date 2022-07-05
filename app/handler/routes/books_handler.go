package routes

import (
	"bookstore/app/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BooksHandler struct {
	model *model.Books
}

func (h *BooksHandler) GetAll(c echo.Context) error {
	return c.String(http.StatusOK, "/users/:id")
}

func (h *BooksHandler) GetById(c echo.Context) error {
	return nil
}

func (h *BooksHandler) Create(c echo.Context) error {
	return nil
}

func (h *BooksHandler) Update(c echo.Context) error {
	return nil
}

func (h *BooksHandler) Delete(c echo.Context) error {
	return nil
}
