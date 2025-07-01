package middlewares

import (
	"github.com/NurilH/belajar-gin-gonic/pkg/common/constants"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {

	auth := ctx.GetHeader("client-key")

	if auth != constants.ClientKey {
		ctx.AbortWithStatusJSON(401, gin.H{
			"message": "Token tidak sesuai",
		})
	}
}
