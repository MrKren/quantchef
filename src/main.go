package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/contrib/static"
)

func main() {
    router := gin.Default()
    router.Use(static.Serve("/", static.LocalFile("../frontend/build", true)))
    router.StaticFile("/logo", "../frontend/src/logo.svg")

    router.Run("0.0.0.0:8080")
}
