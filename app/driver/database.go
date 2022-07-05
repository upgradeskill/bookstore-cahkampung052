package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	MYSQL_USERNAME string
	MYSQL_PASSWORD string
	MYSQL_HOST     string
	MYSQL_DB       string
	MYSQL_PORT     string
	db             *gorm.DB
}

func (d *DB) Connect() (*gorm.DB, error) {
	dsn := d.MYSQL_USERNAME + ":" + d.MYSQL_PASSWORD + "@tcp(" + d.MYSQL_HOST + ":" + d.MYSQL_PORT + ")/" + d.MYSQL_DB + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to database")
	}

	d.db = db
	return d.db, nil
}

func (d *DB) Get() *DB {
	return d
}
