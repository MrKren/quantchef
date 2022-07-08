package recipies

import (
	"net/http"

	"github.com/MrKren/quantchef/src/models"
	"github.com/gin-gonic/gin"
)

func (h handler) PostRecipie(c *gin.Context) {
	body := models.Recipie{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var recipie models.Recipie
	recipie.Title = body.Title
	recipie.Author = body.Author
	recipie.Ingredients = body.Ingredients
	recipie.Instructions = body.Instructions

	if result := h.DB.Create(&recipie); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &recipie)
}