package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func Init() *gorm.DB {
	var err error
	dsn := "monty:password@tcp(127.0.0.1:3306)/electro?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
