package postgres

import (
	"context"

	"github.com/NurilH/belajar-gin-gonic/model"
	"github.com/NurilH/belajar-gin-gonic/module/users"
	"gorm.io/gorm"
)

type usersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) users.UsersRepository {
	return usersRepository{
		db: db,
	}
}

func (r usersRepository) GetAllUsers(ctx context.Context) (user []model.User, err error) {
	err = r.db.WithContext(ctx).Find(&user).Error
	return
}

func (r usersRepository) GetUserByEmail(ctx context.Context, email string) (user model.User, err error) {
	err = r.db.First(&user).Where("email=?", email).Error
	return
}
