package repository

import (
	"bookstore/http/model"
)

type BooksRepo interface {
	Fetch() ([]model.Books, error)
	GetById(id int32) (model.Books, error)
	Create(b *model.Books) (model.Books, error)
	Update(b *model.Books) (model.Books, error)
	Delete(id int32) (bool, error)
}
