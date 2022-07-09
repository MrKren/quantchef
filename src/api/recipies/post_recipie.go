package recipies

import (
	"log"
	"net/http"

	"github.com/MrKren/quantchef/src/models"
	"github.com/gin-gonic/gin"
)

func (h handler) PostRecipie(c *gin.Context) {
	body := models.RecipieRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		gin_err := c.AbortWithError(http.StatusBadRequest, err)
		if gin_err != nil {
			log.Fatalln(gin_err)
		}
		return
	}

	var recipie models.Recipie
	recipie.Title = body.Title
	recipie.Author = body.Author
	recipie.Ingredients = body.Ingredients
	recipie.Instructions = body.Instructions

	if result := h.DB.Create(&recipie); result.Error != nil {
		gin_err := c.AbortWithError(http.StatusNotFound, result.Error)
		if gin_err != nil {
			log.Fatalln(gin_err)
		}
		return
	}

	c.JSON(http.StatusCreated, &recipie)
}
