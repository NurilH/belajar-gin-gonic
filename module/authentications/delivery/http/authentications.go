package http

import (
	"net/http"

	"github.com/NurilH/belajar-gin-gonic/model"
	"github.com/NurilH/belajar-gin-gonic/module/authentications"
	"github.com/gin-gonic/gin"
)

type AuthHTTPDelivery struct {
	route   *gin.RouterGroup
	service authentications.AuthenticationsService
}

func AuthNewDelivery(route *gin.RouterGroup, service authentications.AuthenticationsService) (routeGroup *gin.RouterGroup) {
	authHTTPDelivery := AuthHTTPDelivery{
		route:   route,
		service: service,
	}

	routeGroup = route.Group("/")
	{
		routeGroup.POST("login", authHTTPDelivery.Login)
		routeGroup.POST("signup", authHTTPDelivery.SignUp)
	}

	return
}
func (a AuthHTTPDelivery) SignUp(c *gin.Context) {
	var req model.SignUpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
			"erorr":   err.Error(),
		})
		return
	}

	err := a.service.SignUp(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error insert data",
			"erorr":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": req,
	})

}

func (a AuthHTTPDelivery) Login(c *gin.Context) {

	var req model.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := a.service.Login(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.AbortWithStatusJSON(200, gin.H{
		"message":  "berhasil login",
		"response": result,
	})

}
