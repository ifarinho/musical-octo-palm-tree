package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"mail-app/api/models"
	"mail-app/config"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func Init() *gorm.DB {
	var err error

	db, err = gorm.Open(mysql.Open(config.Dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.Company{}, &models.Mail{}, &models.User{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
