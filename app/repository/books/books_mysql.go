package books

import (
	"bookstore/app/model"

	"gorm.io/gorm"
)

type MysqlBooksRepo struct {
	c *gorm.DB
}

// func MysqlBooksRepo() repository.BooksRepo {
// 	// db := &gorm.DB{}
// 	// booksRepo := MysqlBooksRepo{}

// 	// booksRepo.connect(db)
// 	return &MysqlBooksRepo{}
// }

func (conn *MysqlBooksRepo) Connect(c *gorm.DB) {
	conn.c = c
}

func (conn *MysqlBooksRepo) Fetch() ([]*model.Books, error) {
	books := model.Books{}
	result := conn.c.Find(&books)

	if result.Error != nil {
		return nil, result.Error
	}

	return nil, nil
}

func (conn *MysqlBooksRepo) GetById(bookId int32) (*model.Books, error) {
	return nil, nil
}

func (conn *MysqlBooksRepo) Create(books *model.Books) (*model.Books, error) {
	return nil, nil
}

func (conn *MysqlBooksRepo) Update(books *model.Books) (*model.Books, error) {
	return nil, nil
}

func (conn *MysqlBooksRepo) Delete(bookId int32) (bool, error) {
	return false, nil
}
