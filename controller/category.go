package controller

import (
	"game-api/databases"
	"game-api/utils/logic"
	"game-api/utils/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCategory(c *gin.Context) {
	var result gin.H

	category, err := logic.GetAllCategory(databases.DBConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": category,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertCategory(c *gin.Context) {
	var Category structs.Category

	err := c.ShouldBindJSON(&Category)

	if err != nil {
		panic(err)
	}

	err = logic.InsertCategory(databases.DBConnection, Category)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Berhasil Menambahkan Data",
	})
}

func UpdateCategory(c *gin.Context) {
	var Category structs.Category

	id := c.Param("id")

	err := c.ShouldBindJSON(&Category)

	if err != nil {
		panic(err)
	}

	Category.Id = id

	err = logic.UpdateCategory(databases.DBConnection, Category)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Berhasil Menupdate Data",
	})
}

func DeleteCategory(c *gin.Context) {
	var Category structs.Category

	id := c.Param("id")

	Category.Id = id

	err := logic.DeletedCategory(databases.DBConnection, Category)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Berhasil Menghapus Data",
	})
}
