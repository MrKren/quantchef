package recipies

import (
	"log"
	"net/http"

	"github.com/MrKren/quantchef/src/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetRecipies(c *gin.Context) {
	var recipies []models.Recipie

	result := h.DB.Preload(
		"Ingredients",
	).Preload(
		"Instructions",
	).Preload(
		"Instructions.Steps",
	).Find(&recipies)
	if result.Error != nil {
		gin_err := c.AbortWithError(http.StatusNotFound, result.Error)
		if gin_err != nil {
			log.Fatalln(gin_err)
		}
		return
	}

	c.JSON(http.StatusOK, &recipies)
}
