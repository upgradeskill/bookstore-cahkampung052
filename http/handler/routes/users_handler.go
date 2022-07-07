package routes

import (
	"bookstore/app/pkg"
	domain "bookstore/http/domain/users"
	"bookstore/http/model"
	"bookstore/http/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
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
	var roles map[string]map[string]bool

	for _, val := range users {

		// Set custom claims
		claims := &pkg.JwtCustomClaims{
			val.Id,
			val.Email,
			val.Roles,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}

		// Decode json roles
		err := json.Unmarshal([]byte(val.Roles), &roles)
		if err != nil {
			fmt.Println(err)
		}

		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			fmt.Println(err)
		}

		respUsers = append(respUsers, model.UserPayload{
			Name:  val.Name,
			Email: val.Email,
			Id:    val.Id,
			Roles: roles,
			Token: t,
		})
	}

	return c.JSON(http.StatusOK, pkg.ResponseFormatter("success get all data", respUsers))
}
