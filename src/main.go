package main

import (
	"log"

	"github.com/MrKren/quantchef/src/db"
	"github.com/MrKren/quantchef/src/api/recipies"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("../frontend/build", true)))
	router.StaticFile("/logo", "../frontend/src/logo.svg")

	h := db.Init()

	recipies.RegisterRoutes(router, h)

	err := router.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatalln(err)
	}
}
