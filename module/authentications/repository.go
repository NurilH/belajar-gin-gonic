package authentications

import (
	"github.com/NurilH/belajar-gin-gonic/model"
	"github.com/gin-gonic/gin"
)

type AuthRepository interface {
	SignUp(c *gin.Context, req model.SignUpRequest) error
}

type AuthRedisRepository interface {
	Save(ctx *gin.Context, redisKey string, value string) error
	GetKey(ctx *gin.Context, redisKey string) (res *model.LoginRespons, err error)
}
