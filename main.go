package main

import (
	"fmt"

	users "github.com/NurilH/belajar-gin-gonic/module/users/delivery/http"
	"github.com/NurilH/belajar-gin-gonic/pkg/config"
	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadEnv("config.env")
	conf := config.NewConfig()

	router := gin.Default()

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			InitModuleUsers(v1)
		}
	}

	router.Run(fmt.Sprintf(":%s", conf.AppPort))
}

func InitModuleUsers(router *gin.RouterGroup) *gin.RouterGroup {

	return users.UsersNewDelivery(router)
}
