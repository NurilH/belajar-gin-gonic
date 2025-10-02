package main

import (
	"fmt"
	"log"

	auth "github.com/NurilH/belajar-gin-gonic/module/authentications/delivery/http"
	authRepository "github.com/NurilH/belajar-gin-gonic/module/authentications/repository/postgres"
	authRedisRepository "github.com/NurilH/belajar-gin-gonic/module/authentications/repository/redis"
	authService "github.com/NurilH/belajar-gin-gonic/module/authentications/service"
	document "github.com/NurilH/belajar-gin-gonic/module/documents/delivery/http"
	users "github.com/NurilH/belajar-gin-gonic/module/users/delivery/http"
	usersRepository "github.com/NurilH/belajar-gin-gonic/module/users/repository/postgres"
	usersService "github.com/NurilH/belajar-gin-gonic/module/users/service"
	"github.com/NurilH/belajar-gin-gonic/pkg/common/middlewares"
	"github.com/NurilH/belajar-gin-gonic/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func main() {
	config.LoadEnv("config.env")
	conf := config.NewConfig()

	db, err := config.NewDBGormV2(conf)
	if err != nil {
		log.Fatal(err)
	}

	rdb, err := config.NewRedisConnection(&conf.Redis)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.Static("/static", "./static")

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		v1.Use(middlewares.AuthMiddleware)
		{
			InitModuleDocuments(v1, db)
			InitModuleAuth(v1, db, rdb)
			InitModuleUsers(v1, db)

		}
	}

	go func() {
		router.Run(fmt.Sprintf(":%s", conf.AppPort))
	}()

	router.RunTLS(":443", "/etc/ssl/certs/server.crt", "/etc/ssl/certs/server.key")
}

func InitModuleDocuments(router *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {
	// userRepo := usersRepository.NewUsersRepository(db)
	// authRepo := authRepository.NewAuthRepository(db)
	// authSvc := authService.NewAuthService(authRepo, userRepo)
	return document.DocumentsNewDelivery(router)
}
func InitModuleAuth(router *gin.RouterGroup, db *gorm.DB, rdb *redis.Client) *gin.RouterGroup {
	userRepo := usersRepository.NewUsersRepository(db)
	authRepo := authRepository.NewAuthRepository(db)
	authRedisRepo := authRedisRepository.NewAuthRedisRepository(rdb)
	authSvc := authService.NewAuthService(authRedisRepo, authRepo, userRepo)
	return auth.AuthNewDelivery(router, authSvc)
}

func InitModuleUsers(router *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {
	userRepo := usersRepository.NewUsersRepository(db)
	userSvc := usersService.NewUsersService(userRepo)

	return users.UsersNewDelivery(router, userSvc)
}
