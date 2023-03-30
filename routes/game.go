package routes

import (
	"game-api/controller"

	"github.com/gin-gonic/gin"
)

func GameRoute(router *gin.Engine) {
	game := router.Group("/game")
	{
		game.GET("/:id", controller.GetAllGame)
		game.POST("/", controller.InsertGame)
		game.PUT("/:id", controller.UpdateGame)
		game.DELETE("/:id", controller.DeleteGame)
	}
}
