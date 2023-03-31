package routes

import (
	"game-api/controller"
	"game-api/middleware"

	"github.com/gin-gonic/gin"
)

func CommentRoute(router *gin.Engine) {
	comment := router.Group("/comment").Use(middleware.VerifyJWT())
	{
		comment.POST("/", controller.InsertComment)
		comment.PUT("/:id", controller.UpdateComment)
		comment.DELETE("/:id", controller.DeleteComment)
	}
}
