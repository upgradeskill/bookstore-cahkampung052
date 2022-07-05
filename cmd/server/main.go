package main

import (
	database "bookstore/app/driver"
	"bookstore/app/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db := database.DB{
		MYSQL_USERNAME: "root",
		MYSQL_PASSWORD: "bismillah",
		MYSQL_HOST:     "localhost",
		MYSQL_DB:       "bookstore",
		MYSQL_PORT:     "3306",
	}

	// Connect to Mysql
	db.Connect()

	// migration default data
	db.ImportDummyData()

	// Create Endpoint Routing
	handler.SetUpRoutes(e)

	// Start server
	e.Start(":1323")
}
