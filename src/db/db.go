package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	// TODO: add environment variable manager for production etc.
	// db:3306 is the db address on the docker-compose network
	db_url := "root:root@tcp(db:3306)/test_db"

	db, err := gorm.Open(mysql.Open(db_url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
