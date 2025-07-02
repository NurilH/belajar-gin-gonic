package service

import (
	"os"
	"time"

	"github.com/NurilH/belajar-gin-gonic/model"
	"github.com/NurilH/belajar-gin-gonic/module/authentications"
	"github.com/NurilH/belajar-gin-gonic/module/users"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type authenticationsService struct {
	repoAuth authentications.AuthRepository
	repoUser users.UsersRepository
}

func NewAuthService(repoAuth authentications.AuthRepository, repoUser users.UsersRepository) authentications.AuthenticationsService {
	return &authenticationsService{
		repoAuth: repoAuth,
		repoUser: repoUser,
	}
}

func (a authenticationsService) SignUp(c *gin.Context, req model.SignUpRequest) (err error) {

	hashPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	req.Password = string(hashPass)
	err = a.repoAuth.SignUp(c, req)
	return
}

func (a authenticationsService) Login(c *gin.Context, req model.LoginRequest) (result model.LoginRespons, err error) {
	user, err := a.repoUser.GetUserByEmail(c, req.Email)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return
	}

	// Create the Claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"use_id": user.ID,
		"email":  user.Email,
		"exp":    time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return
	}

	result.Token = &tokenString
	return
}
