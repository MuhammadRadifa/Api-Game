package controller

import (
	"game-api/databases"
	"game-api/utils/logic"
	"game-api/utils/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var Users structs.Users

	err := c.ShouldBindJSON(&Users)

	if err != nil {
		panic(err)
	}

	err = logic.Register(databases.DBConnection, Users)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Berhasil Registrasi",
	})
}
