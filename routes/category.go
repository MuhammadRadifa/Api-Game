package routes

import (
	"game-api/controller"

	"github.com/gin-gonic/gin"
)

func CategoryRoute(router *gin.Engine) {
	Category := router.Group("/category")
	{
		Category.GET("/", controller.GetAllCategory)
		Category.POST("/", controller.InsertCategory)
		Category.PUT("/:id", controller.UpdateCategory)
		Category.DELETE("/:id", controller.DeleteCategory)
	}
}
