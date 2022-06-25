package main

import (
	"fmt"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("../frontend/build", true)))
	router.StaticFile("/logo", "../frontend/src/logo.svg")

	err := router.Run("0.0.0.0:8080")
	if err != nil {
		fmt.Print(err)
	}
}
