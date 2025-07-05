package common

import (
	"github.com/NurilH/belajar-gin-gonic/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Controller struct {
}

type UserInfo struct {
	UserID float64 `json:"user_id"`
	Email  string  `json:"email"`
}

func (c *Controller) User(ctx *gin.Context) (result *UserInfo) {
	tokenStr, err := ctx.Cookie("Authorization")
	if err != nil || tokenStr == "" {
		tokenStr = ctx.GetHeader("Authorization")
		if tokenStr == "" {
			return
		}
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Env("SECRET_KEY")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		result = new(UserInfo)
		result.Email = claims["email"].(string)
		result.UserID = claims["user_id"].(float64)
		return
	} else {
		return
	}
}
