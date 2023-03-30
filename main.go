package main

import (
	"game-api/databases"
	"game-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	databases.Connection()

	router := gin.Default()
	routes.UserRoute(router)
	routes.CategoryRoute(router)
	routes.GameRoute(router)
	routes.CommentRoute(router)
	routes.RatingRoute(router)

	router.Run("localhost:8080")
}
