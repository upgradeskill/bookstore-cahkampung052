package routes

import (
	"bookstore/app/pkg"
	domain "bookstore/http/domain/users"
	"bookstore/http/model"
	"bookstore/http/repository"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Users struct {
	Repo repository.UserRepo
}

func UsersHandler(db *gorm.DB) *Users {
	return &Users{
		Repo: domain.MysqlUserDomain(db),
	}
}

func (u *Users) GetAll(c echo.Context) error {
	// Get all users and validate if no error
	users, err := u.Repo.Fetch()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pkg.ResponseFormatter(err.Error(), nil))
	}

	var respUsers []model.UserPayload
	var roles map[string]interface{}

	for _, val := range users {

		err := json.Unmarshal([]byte(val.Roles), &roles)
		if err != nil {
			fmt.Println(err)
		}

		respUsers = append(respUsers, model.UserPayload{
			Name:  val.Name,
			Email: val.Email,
			Id:    val.Id,
			Roles: roles,
		})
	}

	return c.JSON(http.StatusOK, pkg.ResponseFormatter("success get all data", respUsers))
}
