package http

import (
	"github.com/gin-gonic/gin"
)

type UsersHTTPDelivery struct {
	route *gin.RouterGroup
}

func UsersNewDelivery(route *gin.RouterGroup) (routeGroup *gin.RouterGroup) {
	usersHTTPDelivery := UsersHTTPDelivery{
		route: route,
	}

	routeGroup = route.Group("/user")
	{
		routeGroup.GET("", usersHTTPDelivery.GetAllUsers)
		routeGroup.GET("/detail", usersHTTPDelivery.GetDetailUser)
	}

	return
}

func (u UsersHTTPDelivery) GetAllUsers(ctx *gin.Context) {
	error := false
	if error {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "bad request",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "All Users",
	})

}

func (u UsersHTTPDelivery) GetDetailUser(ctx *gin.Context) {
	error := false
	if error {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "bad request",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Detail User",
	})

}
