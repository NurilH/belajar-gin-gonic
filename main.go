package main

import (
	"fmt"
	"log"

	auth "github.com/NurilH/belajar-gin-gonic/module/authentications/delivery/http"
	users "github.com/NurilH/belajar-gin-gonic/module/users/delivery/http"
	usersRepository "github.com/NurilH/belajar-gin-gonic/module/users/repository/postgres"
	usersService "github.com/NurilH/belajar-gin-gonic/module/users/service"
	"github.com/NurilH/belajar-gin-gonic/pkg/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {

	config.LoadEnv("config.env")
	conf := config.NewConfig()

	router := gin.Default()

	db, err := config.NewDBGormV2(conf)
	if err != nil {
		log.Fatal(err)
	}

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			InitModuleUsers(v1, db)

			InitModuleAuth(v1)

		}
	}

	router.Run(fmt.Sprintf(":%s", conf.AppPort))
}

func InitModuleAuth(router *gin.RouterGroup) *gin.RouterGroup {

	return auth.AuthNewDelivery(router)
}

func InitModuleUsers(router *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {
	userRepo := usersRepository.NewUsersRepository(db)
	userSvc := usersService.NewUsersService(userRepo)

	return users.UsersNewDelivery(router, userSvc)
}
