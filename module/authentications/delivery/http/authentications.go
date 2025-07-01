package http

import (
	"net/http"

	"github.com/NurilH/belajar-gin-gonic/model"
	"github.com/gin-gonic/gin"
)

type AuthHTTPDelivery struct {
	route *gin.RouterGroup
}

func AuthNewDelivery(route *gin.RouterGroup) (routeGroup *gin.RouterGroup) {
	authHTTPDelivery := AuthHTTPDelivery{
		route: route,
	}

	routeGroup = route.Group("/")
	{
		routeGroup.POST("login", authHTTPDelivery.AuthLogin)
	}

	return
}

func (a AuthHTTPDelivery) AuthLogin(ctx *gin.Context) {

	var req model.RequestLogin

	// if err := ctx.Bind(&req); err != nil {
	// 	ctx.AbortWithStatusJSON(400, gin.H{
	// 		"message": "erro login",
	// 	})
	// 	return
	// }

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// if err := ctx.Validate(&req); err != nil {
	// 	ctx.AbortWithStatusJSON(400, gin.H{
	// 		"message": "erro login",
	// 	})
	// 	return
	// }

	ctx.AbortWithStatusJSON(200, gin.H{
		"message": "berhasil login",
		"payload": req,
	})

}
