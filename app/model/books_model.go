package model

import (
	"time"

	"gorm.io/gorm"
)

type Books struct {
	Id          int    `gorm:"primaryKey;autoIncrement"`
	Title       string `gorm:"size:150;unique:true"`
	Price       int
	Author      string `gorm:"size:150"`
	Publisher   string `gorm:"size:150"`
	PublishDate string `gorm:"type:date"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
