package controller

import (
	"game-api/databases"
	"game-api/middleware"
	"game-api/utils/logic"
	"game-api/utils/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InsertComment(c *gin.Context) {

	var Comment structs.Comment

	err := c.ShouldBindJSON(&Comment)

	if err != nil {
		panic(err)
	}

	Comment.Users_id, _ = middleware.ExtractClaims(c, "id")

	if err != nil {
		panic(err)
	}

	err = logic.InsertComment(databases.DBConnection, Comment)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Berhasil Menambahkan Data",
	})
}

func UpdateComment(c *gin.Context) {
	var Comment structs.Comment

	id := c.Param("id")

	err := c.ShouldBindJSON(&Comment)

	if err != nil {
		panic(err)
	}

	Comment.Id = id

	err = logic.UpdateComment(databases.DBConnection, Comment)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Berhasil Menupdate Data",
	})
}

func DeleteComment(c *gin.Context) {
	var Comment structs.Comment

	id := c.Param("id")

	Comment.Id = id

	err := logic.DeletedComment(databases.DBConnection, Comment)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Berhasil Menghapus Data",
	})
}
