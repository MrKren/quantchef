package recipies

import (
	"log"
	"net/http"

	"github.com/MrKren/quantchef/src/models"
	"github.com/gin-gonic/gin"
)

func (h handler) DeleteRecipie(c *gin.Context) {
	ID := c.Param("ID")

	var recipie models.Recipie

	result := h.DB.First(&recipie, ID)
	if result.Error != nil {
		gin_err := c.AbortWithError(http.StatusNotFound, result.Error)
		if gin_err != nil {
			log.Fatalln(gin_err)
		}
		return
	}

	h.DB.Delete(&recipie)

	c.Status(http.StatusOK)
}
