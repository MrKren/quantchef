package db

import (
	"log"

	"github.com/MrKren/quantchef/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	// TODO: add environment variable manager for production etc.
	// db:3306 is the db address on the docker-compose network
	db_url := "root:root@tcp(db:3306)/quantchef"

	db, err := gorm.Open(mysql.Open(db_url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Recipie{})
	db.AutoMigrate(&models.Ingredient{})
	db.AutoMigrate(&models.Dish{})
	db.AutoMigrate(&models.Step{})

	return db
}
