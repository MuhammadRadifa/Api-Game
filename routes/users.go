package routes

import (
	"game-api/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {

	router.POST("/register", controller.Register)

	router.POST("/login", controller.Login)
}
