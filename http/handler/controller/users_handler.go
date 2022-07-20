package handler

import (
	"bookstore/app/pkg"
	"bookstore/http/domain"
	"bookstore/http/model"
	repository "bookstore/http/repository/users"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type userHandler struct {
	Repo repository.UserRepo
}

func UsersHandler(db *gorm.DB) *userHandler {
	return &userHandler{
		Repo: domain.MysqlUserDomain(db),
	}
}

func (u *userHandler) GetAll(c echo.Context) error {
	// Get all users and validate if no error
	users, err := u.Repo.Fetch()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pkg.ResponseFormatter(err.Error(), nil))
	}

	var respUsers []model.UserPayload

	for _, val := range users {
		var roles map[string]map[string]bool

		// Decode json roles
		err := json.Unmarshal([]byte(val.Roles), &roles)
		if err != nil {
			fmt.Println(err)
		}

		// Set custom claims
		claims := &model.JwtUserClaims{
			val.Id,
			val.Email,
			val.Roles,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
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
