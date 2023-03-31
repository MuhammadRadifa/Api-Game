package routes

import (
	"game-api/controller"
	"game-api/middleware"

	"github.com/gin-gonic/gin"
)

func CategoryRoute(router *gin.Engine) {
	Category := router.Group("/category").Use(middleware.VerifyJWT())
	{
		Category.GET("/", controller.GetAllCategory)
		Category.POST("/", controller.InsertCategory).Use(middleware.IsAdmin())
		Category.PUT("/:id", controller.UpdateCategory).Use(middleware.IsAdmin())
		Category.DELETE("/:id", controller.DeleteCategory).Use(middleware.IsAdmin())
	}
}
