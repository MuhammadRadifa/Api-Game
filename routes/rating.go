package routes

import (
	"game-api/controller"

	"github.com/gin-gonic/gin"
)

func RatingRoute(router *gin.Engine) {
	rating := router.Group("/rating")
	{
		rating.POST("/", controller.InsertRating)
		rating.PUT("/:id", controller.UpdateRating)
		rating.DELETE("/:id", controller.DeleteRating)
	}
}
