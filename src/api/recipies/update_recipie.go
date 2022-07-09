package recipies

import (
	"log"
	"net/http"

	"github.com/MrKren/quantchef/src/models"
	"github.com/gin-gonic/gin"
)

func (h handler) UpdateRecipie(c *gin.Context) {
	ID := c.Param("ID")
	body := models.RecipieRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		gin_err := c.AbortWithError(http.StatusBadRequest, err)
		if gin_err != nil {
			log.Fatalln(gin_err)
		}
		return
	}

	var recipie models.Recipie

	result := h.DB.First(&recipie, ID)
	if result.Error != nil {
		gin_err := c.AbortWithError(http.StatusNotFound, result.Error)
		if gin_err != nil {
			log.Fatalln(gin_err)
		}
		return
	}

	recipie.Title = body.Title
	recipie.Author = body.Author
	recipie.Ingredients = body.Ingredients
	recipie.Instructions = body.Instructions

	h.DB.Save(&recipie)

	c.JSON(http.StatusOK, &recipie)
}
