package main

import (
	"fmt"
	"log"

	auth "github.com/NurilH/belajar-gin-gonic/module/authentications/delivery/http"
	authRepository "github.com/NurilH/belajar-gin-gonic/module/authentications/repository/postgres"
	authService "github.com/NurilH/belajar-gin-gonic/module/authentications/service"
	document "github.com/NurilH/belajar-gin-gonic/module/documents/delivery/http"
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

	router.Static("/static", "./static")

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		v1.Use(middlewares.AuthMiddleware)
		{
			InitModuleDocuments(v1, db)
			InitModuleAuth(v1, db)
			InitModuleUsers(v1, db)

		}
	}

	router.Run(fmt.Sprintf(":%s", conf.AppPort))
}

func InitModuleDocuments(router *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {
	// userRepo := usersRepository.NewUsersRepository(db)
	// authRepo := authRepository.NewAuthRepository(db)
	// authSvc := authService.NewAuthService(authRepo, userRepo)
	return document.DocumentsNewDelivery(router)
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
