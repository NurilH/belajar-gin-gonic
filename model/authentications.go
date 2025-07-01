package model

type RequestLogin struct {
	Username string `json:"username" binding:"required,lte=5"`
	Password string `json:"password" binding:"required"`
	UserID   int    `json:"user_id" binding:"lte=5"`
}
