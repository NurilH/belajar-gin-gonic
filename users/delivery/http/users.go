package http

import (
	"github.com/gin-gonic/gin"
)

func UsersNewDelivery(router *gin.Engine) {
	routeGroup := router.Group("/user")

	routeGroup.GET("", GetAllUsers)
	routeGroup.GET("/detail", GetAllUsers)

	return
}

func GetAllUsers(ctx *gin.Context) {
	error := true
	if error {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "bad request",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "All Users",
	})
	return
}

func GetDetailUser(ctx *gin.Context) {
	error := true
	if error {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "bad request",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Detail User",
	})
	return
}
