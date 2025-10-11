package http

import (
	"strconv"

	"github.com/NurilH/belajar-gin-gonic/module/users"
	"github.com/NurilH/belajar-gin-gonic/pkg/common"
	"github.com/gin-gonic/gin"
)

type UsersHTTPDelivery struct {
	common.Controller
	route        *gin.RouterGroup
	usersService users.UsersService
}

func UsersNewDelivery(route *gin.RouterGroup, usersService users.UsersService) (routeGroup *gin.RouterGroup) {
	usersHTTPDelivery := UsersHTTPDelivery{
		route:        route,
		usersService: usersService,
	}

	routeGroup = route.Group("/user")
	{
		routeGroup.GET("", usersHTTPDelivery.GetAllUsers)
		routeGroup.GET("/detail", usersHTTPDelivery.GetDetailUser)
		routeGroup.GET("/:id", usersHTTPDelivery.GetUserByID)
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

	result, err := u.usersService.GetAllUsers(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "bad request",
			"error":   err.Error(),
		})
		return

	}

	ctx.JSON(200, gin.H{
		"message": "All Users",
		"data":    result,
	})

}

func (u UsersHTTPDelivery) GetDetailUser(ctx *gin.Context) {

	result := u.User(ctx)

	ctx.JSON(200, gin.H{
		"message": "Detail User from token",
		"data":    result,
	})

}
func (u UsersHTTPDelivery) GetUserByID(ctx *gin.Context) {

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "err convert",
			"error":   err.Error(),
		})
		return

	}

	result, err := u.usersService.GetUserByID(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "bad request",
			"error":   err.Error(),
		})
		return

	}

	ctx.JSON(200, gin.H{
		"message": "Detail User By id",
		"data":    result,
	})

}
