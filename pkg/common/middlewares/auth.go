package middlewares

import (
	"os"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {

	auth := ctx.GetHeader("client-key")

	if auth != os.Getenv("SECRET_KEY") {
		ctx.AbortWithStatusJSON(401, gin.H{
			"message": "Token tidak sesuai",
		})
	}
}
