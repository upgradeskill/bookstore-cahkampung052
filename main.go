package main

import (
	"bookstore/app/database"
	"bookstore/app/driver"
	"bookstore/http/handler"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type RequestValidator struct {
	validator *validator.Validate
}

func (cv *RequestValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

var config = driver.ConfigDb{
	MYSQL_USERNAME: "root",
	MYSQL_PASSWORD: "bismillah",
	MYSQL_HOST:     "localhost",
	MYSQL_DB:       "bookstore",
	MYSQL_PORT:     "3306",
}

func main() {
	e := echo.New()
	e.Validator = &RequestValidator{validator: validator.New()}

	connection, _ := driver.ConnectMysql(config)

	handler.SetUpRoutes(e, connection)
	database.RunMigration(connection)

	// Start server
	e.Start(":4200")
}
