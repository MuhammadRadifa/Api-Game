package controller

import (
	"fmt"
	"game-api/databases"
	"game-api/middleware"
	"game-api/utils/logic"
	"game-api/utils/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InsertRating(c *gin.Context) {
	var Rating structs.Rating

	err := c.ShouldBindJSON(&Rating)

	if err != nil {
		panic(err)
	}

	Rating.Users_id, _ = middleware.ExtractClaims(c, "id")

	if err != nil {
		panic(err)
	}

	fmt.Print(Rating.Users_id)

	err = logic.InsertRating(databases.DBConnection, Rating)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Berhasil Menambahkan Data",
	})
}

func UpdateRating(c *gin.Context) {
	var Rating structs.Rating

	id := c.Param("id")

	err := c.ShouldBindJSON(&Rating)

	if err != nil {
		panic(err)
	}

	Rating.Id = id

	err = logic.UpdateRating(databases.DBConnection, Rating)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Berhasil Menupdate Data",
	})
}

func DeleteRating(c *gin.Context) {
	var Rating structs.Rating

	id := c.Param("id")

	Rating.Id = id

	err := logic.DeletedRating(databases.DBConnection, Rating)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Berhasil Menghapus Data",
	})
}
