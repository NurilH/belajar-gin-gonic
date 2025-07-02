package postgres

import (
	"github.com/NurilH/belajar-gin-gonic/model"
	"github.com/NurilH/belajar-gin-gonic/module/authentications"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) authentications.AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (a *authRepository) SignUp(c *gin.Context, req model.SignUpRequest) error {
	err := a.db.Create(&req).Error
	return err
}
