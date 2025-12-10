package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	auth "github.com/NurilH/belajar-gin-gonic/module/authentications/delivery/http"
	authRepository "github.com/NurilH/belajar-gin-gonic/module/authentications/repository/postgres"
	authRedisRepository "github.com/NurilH/belajar-gin-gonic/module/authentications/repository/redis"
	authService "github.com/NurilH/belajar-gin-gonic/module/authentications/service"
	document "github.com/NurilH/belajar-gin-gonic/module/documents/delivery/http"
	users "github.com/NurilH/belajar-gin-gonic/module/users/delivery/http"
	usersRepository "github.com/NurilH/belajar-gin-gonic/module/users/repository/postgres"
	usersService "github.com/NurilH/belajar-gin-gonic/module/users/service"
	"github.com/NurilH/belajar-gin-gonic/pkg/common/helpers"
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

	staticDir := helpers.GetStaticDir()

	router := gin.Default()

	router.Static("/static", staticDir)

	router.GET("/debug/paths", func(c *gin.Context) {
		wd, _ := os.Getwd()
		c.JSON(200, gin.H{
			"working_dir": wd,
			"static_dir":  staticDir,
			"exists":      dirExists(staticDir),
			"files":       listFiles(staticDir),
		})
	})

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

	if strings.ToUpper(conf.AppEnv) == "LOCAL" {
		router.Run(fmt.Sprintf(":%s", conf.AppPort))
	} else {
		go func() {
			router.Run(fmt.Sprintf(":%s", conf.AppPort))
		}()

		router.RunTLS(":443", "/app/certs/server.crt", "/app/certs/server.key")
	}

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

func dirExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func listFiles(dir string) []string {
	files, err := os.ReadDir(dir)
	if err != nil {
		return []string{}
	}

	var fileList []string
	for _, file := range files {
		fileList = append(fileList, file.Name())
	}
	return fileList
}
