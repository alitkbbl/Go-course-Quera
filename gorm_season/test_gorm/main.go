package main

import (
	GetDBPackage "test_gorm/db"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name   string `gorm:"size:100"`
	Email  string `gorm:"uniqueIndex"`
	Age    int
	Orders []Order `gorm:"foreignKey:UserID"`
}

type Order struct {
	gorm.Model
	UserID uint
	Amount float64
	Status string `gorm:"default:'pending'"`
	User   User   `gorm:"foreignKey:UserID"`
}

func main() {
	db := GetDBPackage.GetDB()
	print(db)

}
