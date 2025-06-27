package users

import (
	"context"

	"github.com/NurilH/belajar-gin-gonic/model"
)

type (
	UsersRepository interface {
		GetAllUsers(ctx context.Context) (user model.Users, err error)
	}
)
