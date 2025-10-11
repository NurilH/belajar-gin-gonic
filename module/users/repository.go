package users

import (
	"context"

	"github.com/NurilH/belajar-gin-gonic/model"
)

type (
	UsersRepository interface {
		GetAllUsers(ctx context.Context) (user []model.User, err error)
		GetUserByEmail(ctx context.Context, email string) (user model.User, err error)
		GetUserByID(ctx context.Context, id int) (user model.User, err error)
	}
)
