package model

import "gorm.io/gorm"

type SignUpRequest struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginRespons struct {
	Token *string `json:"token"`
}

func (SignUpRequest) TableName() string {
	return "users"
}
