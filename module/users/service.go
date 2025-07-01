package users

import (
	"github.com/NurilH/belajar-gin-gonic/model"
	"github.com/gin-gonic/gin"
)

type (
	UsersService interface {
		GetAllUsers(ctx *gin.Context) (result []model.Users, err error)
	}
)
