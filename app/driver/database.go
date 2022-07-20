package driver

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var config = map[string]string{
	"MYSQL_USERNAME": "root",
	"MYSQL_PASSWORD": "bismillah",
	"MYSQL_HOST":     "localhost",
	"MYSQL_DB":       "bookstore",
	"MYSQL_PORT":     "3306",
}

type DB struct {
	db *gorm.DB
}

var connection = &DB{}

func ConnectMysql() (*gorm.DB, error) {
	dsn := config["MYSQL_USERNAME"] + ":" + config["MYSQL_PASSWORD"] + "@tcp(" + config["MYSQL_HOST"] + ":" + config["MYSQL_PORT"] + ")/" + config["MYSQL_DB"] + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to database")
	}

	return db, nil
}
