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
	db_url := "root:root@tcp(db:3306)/quantchef?parseTime=true"

	db, err := gorm.Open(mysql.Open(db_url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(
		&models.Recipie{},
		&models.Ingredient{},
		&models.Dish{},
		&models.Step{},
	)
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
