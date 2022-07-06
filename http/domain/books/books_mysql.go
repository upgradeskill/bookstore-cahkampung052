package domain

import (
	"bookstore/http/model"
	"bookstore/http/repository"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type mysqlBooksRepo struct {
	Db *gorm.DB
}

var books []model.Books
var book model.Books

func MysqlBooksRepo(connection *gorm.DB) repository.BooksRepo {
	return &mysqlBooksRepo{
		Db: connection,
	}
}

func (m *mysqlBooksRepo) Fetch() ([]model.Books, error) {
	m.Db.Raw("SELECT * FROM books").Scan(&books)

	return books, nil
}

func (m *mysqlBooksRepo) GetById(bookId int32) (model.Books, error) {
	m.Db.Raw("SELECT * FROM books WHERE id = ?", bookId).Scan(&book)
	if book.Id == 0 {
		return book, errors.New("Book not found")
	}

	fmt.Println(bookId)

	return book, nil
}

func (m *mysqlBooksRepo) Create(book *model.Books) (model.Books, error) {
	result := m.Db.Model(&book).Create(book)
	if result.Error != nil {
		return model.Books{}, result.Error
	}

	return model.Books{}, nil
}

func (m *mysqlBooksRepo) Update(book *model.Books) (model.Books, error) {
	result := m.Db.Model(&book).Update("id", book.Id)
	if result.Error != nil {
		return model.Books{}, result.Error
	}

	return *book, nil
}

func (m *mysqlBooksRepo) Delete(bookId int32) (bool, error) {
	result := m.Db.Delete(&model.Books{}, bookId)
	fmt.Println("123")
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}
