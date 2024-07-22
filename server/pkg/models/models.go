package models

import (
	"github.com/zinitdev/eSale/pkg/configs"
	"gorm.io/gorm"
)

var db *gorm.DB

type BaseModel struct {
	gorm.Model
	Active      bool	`gorm:"default:true"`
}

type Category struct {
	BaseModel
	Name string `gorm:"unique"`
}

func init() {
	configs.Connect()
	db = configs.GetDB()
	db.AutoMigrate(&Category{})

	categories := []Category{
		{Name: "Mobile"},
		{Name: "Laptop"},
		{Name: "Tablet"},
	}
	db.Create(&categories)
}

func GetCategories() []Category{
	var queries []Category
	db.Find(&queries)
	return queries
}