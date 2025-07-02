package main

import (
	"fmt"
	"log"

	auth "github.com/NurilH/belajar-gin-gonic/module/authentications/delivery/http"
	authRepository "github.com/NurilH/belajar-gin-gonic/module/authentications/repository/postgres"
	authService "github.com/NurilH/belajar-gin-gonic/module/authentications/service"
	users "github.com/NurilH/belajar-gin-gonic/module/users/delivery/http"
	usersRepository "github.com/NurilH/belajar-gin-gonic/module/users/repository/postgres"
	usersService "github.com/NurilH/belajar-gin-gonic/module/users/service"
	"github.com/NurilH/belajar-gin-gonic/pkg/common/middlewares"
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
		v1.Use(middlewares.AuthMiddleware)
		{
			InitModuleUsers(v1, db)

			InitModuleAuth(v1, db)

		}
	}

	router.Run(fmt.Sprintf(":%s", conf.AppPort))
}

func InitModuleAuth(router *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {
	userRepo := usersRepository.NewUsersRepository(db)
	authRepo := authRepository.NewAuthRepository(db)
	authSvc := authService.NewAuthService(authRepo, userRepo)
	return auth.AuthNewDelivery(router, authSvc)
}

func InitModuleUsers(router *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {
	userRepo := usersRepository.NewUsersRepository(db)
	userSvc := usersService.NewUsersService(userRepo)

	return users.UsersNewDelivery(router, userSvc)
}
