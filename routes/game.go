package routes

import (
	"game-api/controller"
	"game-api/middleware"

	"github.com/gin-gonic/gin"
)

func GameRoute(router *gin.Engine) {
	game := router.Group("/game").Use(middleware.VerifyJWT())
	{
		game.GET("/:id", controller.GetAllGame)
		game.POST("/", controller.InsertGame)
		game.PUT("/:id", controller.UpdateGame)
		game.DELETE("/:id", controller.DeleteGame)
	}
}
