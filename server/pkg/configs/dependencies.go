package configs

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
    LoadEnv()
    
	dbUser := GetEnv("DB_USER")
    dbPassword := GetEnv("DB_PASSWORD")
    dbName := GetEnv("DB_NAME")
    dbHost := GetEnv("DB_HOST")
    dbPort := GetEnv("DB_PORT")
	
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        dbUser, dbPassword, dbHost, dbPort, dbName)

	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic(err)
    }

    db = d
}

func GetDB() *gorm.DB {
	return db
}