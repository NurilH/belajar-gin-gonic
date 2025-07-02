package authentications

import (
	"github.com/NurilH/belajar-gin-gonic/model"
	"github.com/gin-gonic/gin"
)

type AuthRepository interface {
	SignUp(c *gin.Context, req model.SignUpRequest) error
}
