package routes

import (
	"game-api/controller"

	"github.com/gin-gonic/gin"
)

func CommentRoute(router *gin.Engine) {
	comment := router.Group("/comment")
	{
		comment.POST("/", controller.InsertComment)
		comment.PUT("/:id", controller.UpdateComment)
		comment.DELETE("/:id", controller.DeleteComment)
	}
}
