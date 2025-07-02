package authentications

import (
	"github.com/NurilH/belajar-gin-gonic/model"
	"github.com/gin-gonic/gin"
)

type AuthenticationsService interface {
	SignUp(c *gin.Context, req model.SignUpRequest) (err error)
	Login(c *gin.Context, req model.LoginRequest) (result model.LoginRespons, err error)
}
