package controller

import (
	"game-api/databases"
	"game-api/utils/logic"
	"game-api/utils/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllGame(c *gin.Context) {
	var result gin.H

	param, _ := strconv.Atoi(c.Param("id"))

	Data, err := logic.GetAllGame(databases.DBConnection, param)

	if err != nil {
		result = gin.H{
			"error": err.Error(),
		}
	} else {
		result = gin.H{
			"result": Data,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertGame(c *gin.Context) {
	var Game structs.Game

	err := c.ShouldBindJSON(&Game)

	if err != nil {
		panic(err)
	}

	err = logic.InsertGame(databases.DBConnection, Game)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Berhasil Menambahkan Data",
	})
}

func UpdateGame(c *gin.Context) {
	var Game structs.Game

	id := c.Param("id")

	err := c.ShouldBindJSON(&Game)

	if err != nil {
		panic(err)
	}

	Game.Id = id

	err = logic.UpdateGame(databases.DBConnection, Game)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Berhasil Menupdate Data",
	})
}

func DeleteGame(c *gin.Context) {
	var Game structs.Game

	id := c.Param("id")

	Game.Id = id

	err := logic.DeletedGame(databases.DBConnection, Game)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Berhasil Menghapus Data",
	})
}
