package main

import (
	"fmt"
	"game-api/databases"
	"game-api/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("config/.env")

	if err != nil {
		fmt.Println("failed Load file environment")
	}

	databases.Connection()

	router := gin.Default()
	routes.UserRoute(router)
	routes.CategoryRoute(router)
	routes.GameRoute(router)
	routes.CommentRoute(router)
	routes.RatingRoute(router)

	router.Run(":" + os.Getenv("PGPORT"))
}
