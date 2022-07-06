package routes

import (
	"bookstore/app/pkg"
	domain "bookstore/http/domain/books"
	"bookstore/http/model"
	"bookstore/http/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Books struct {
	Repo repository.BooksRepo
}

func BooksHandler(db *gorm.DB) *Books {
	return &Books{
		Repo: domain.MysqlBooksDomain(db),
	}
}

func (b *Books) GetAll(c echo.Context) error {
	// Get all book and validate if no error
	books, err := b.Repo.Fetch()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pkg.ResponseFormatter(err.Error(), nil))
	}

	return c.JSON(http.StatusOK, pkg.ResponseFormatter("success get all data", books))
}

func (b *Books) GetById(c echo.Context) error {
	// Validate params must number
	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseFormatter("failed get detail book, please insert number on param", nil))
	}

	// Get detail book and validate if exists
	book, err := b.Repo.GetById(int32(bookId))
	if err != nil {
		return c.JSON(http.StatusNotFound, pkg.ResponseFormatter(err.Error(), nil))
	}

	return c.JSON(http.StatusOK, pkg.ResponseFormatter("success get detail book", book))
}

func (b *Books) Create(c echo.Context) (err error) {
	payload := new(model.Books)
	if err = c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseFormatter(err.Error(), nil))
	}

	if err = c.Validate(payload); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Insert book and validate if no error
	book, err := b.Repo.Create(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseFormatter(err.Error(), nil))
	}

	return c.JSON(http.StatusOK, pkg.ResponseFormatter("success create new book", book))
}

func (b *Books) Update(c echo.Context) (err error) {
	payload := new(model.Books)
	if err = c.Bind(payload); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ResponseFormatter(err.Error(), nil))
	}

	if err = c.Validate(payload); err != nil {
		c.JSON(http.StatusBadRequest, pkg.ResponseFormatter(err.Error(), nil))
	}

	// Insert book and validate if no error
	book, err := b.Repo.Update(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseFormatter(err.Error(), nil))
	}

	return c.JSON(http.StatusOK, pkg.ResponseFormatter("success update book", book))
}

func (b *Books) Delete(c echo.Context) error {
	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseFormatter("failed delete detail book, please insert number on param", nil))
	}

	// Insert book and validate if no error
	book, err := b.Repo.Delete(int32(bookId))
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseFormatter(err.Error(), nil))
	}

	return c.JSON(http.StatusOK, pkg.ResponseFormatter("success delete book", book))
}
