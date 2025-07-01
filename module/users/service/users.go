package service

import (
	"github.com/NurilH/belajar-gin-gonic/model"
	"github.com/NurilH/belajar-gin-gonic/module/users"
	"github.com/gin-gonic/gin"
)

type usersService struct {
	usersRepository users.UsersRepository
}

func NewUsersService(usersRepository users.UsersRepository) users.UsersService {
	return usersService{
		usersRepository: usersRepository,
	}
}

func (s usersService) GetAllUsers(ctx *gin.Context) (result []model.Users, err error) {
	result, err = s.usersRepository.GetAllUsers(ctx.Request.Context())
	return
}
