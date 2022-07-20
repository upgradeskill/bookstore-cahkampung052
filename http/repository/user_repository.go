package repository

import "bookstore/http/model"

type UserRepo interface {
	Fetch() ([]model.User, error)
}
