package middlewares

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/NurilH/belajar-gin-gonic/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(ctx *gin.Context) {

	auth := ctx.GetHeader("client-key")

	if auth != os.Getenv("SECRET_KEY") {
		ctx.AbortWithStatusJSON(401, gin.H{
			"message": "Key tidak sesuai",
		})
	}

	if !Skipper(ctx) {
		ValidateJWT(ctx)
	}
}

func Skipper(ctx *gin.Context) bool {
	path := ctx.FullPath()
	apiV1 := "/api/v1"
	if (ctx.Request.Method == http.MethodPost && strings.EqualFold(path, apiV1+"/login")) ||
		(ctx.Request.Method == http.MethodGet && strings.EqualFold(path, apiV1+"/user")) ||
		(ctx.Request.Method == http.MethodPost && strings.EqualFold(path, apiV1+"/signup")) {
		return true
	}
	return false
}

func ValidateJWT(ctx *gin.Context) {
	tokenStr, err := ctx.Cookie("Authorization")
	if err != nil || tokenStr == "" {
		tokenStr = ctx.GetHeader("Authorization")
		if tokenStr == "" {
			ctx.AbortWithStatusJSON(404, gin.H{
				"message": "Token Not Found",
			})
			return
		}
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Env("SECRET_KEY")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		ctx.AbortWithStatusJSON(401, gin.H{
			"message": "Invalid Token",
			"error":   err.Error(),
		})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		if claims["exp"].(float64) < float64(time.Now().Unix()) {
			ctx.AbortWithStatusJSON(401, gin.H{
				"message": "Expired Token",
			})
			return
		}
		if claims["user_id"] == 0 {
			ctx.AbortWithStatusJSON(401, gin.H{
				"message": "Invalid Token",
			})
			return
		}
	} else {
		ctx.AbortWithStatusJSON(401, gin.H{
			"message": "Invalid Token",
		})
		return
	}
}
