package recipies

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := router.Group("/recipies")
	routes.POST("", h.PostRecipie)
	routes.GET("", h.GetRecipies)
	routes.GET("/:ID", h.GetRecipie)
	routes.PUT("/:ID", h.UpdateRecipie)
	routes.DELETE("/:ID", h.DeleteRecipie)
}
