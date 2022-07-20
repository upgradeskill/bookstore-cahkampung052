package database

import (
	"bookstore/http/model"
	"bookstore/http/repository"
	"fmt"

	"gorm.io/gorm"
)

type userMigration struct {
	conn *gorm.DB
}

func UserMigration(conn *gorm.DB) repository.UserMigrationRepo {
	return &userMigration{
		conn: conn,
	}
}

func (user *userMigration) ImportSeeder() {
	// Skip migration if users table already exist
	// and run migration if users table not exist
	fmt.Println("23")
	isExists := user.conn.Migrator().HasTable("users")
	if isExists {
		return
	}

	// Prepare data for data dummy
	var users = []model.User{
		{
			Name:  "Wahyu",
			Email: "wahyuagung26@gmail.com",
			Roles: "{\"books\":{\"create\":true,\"update\":true,\"delete\":true,\"view\":true}}",
		},
		{
			Name:  "Agung",
			Email: "wahyu.agung@majoo.id",
			Roles: "{\"books\":{\"create\":false,\"update\":true,\"delete\":false,\"view\":true}}",
		}}

	// Create table users and insert batch data dummy
	user.conn.AutoMigrate(model.User{})
	user.conn.Create(users)
}
