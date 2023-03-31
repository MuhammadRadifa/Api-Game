package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, _ := ExtractClaims(ctx, "role")
		if token == "admin" {
			ctx.Next()
			return
		} else {

			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Anda Bukan Admin",
			})
			ctx.Abort()

			return
		}

	}
}
