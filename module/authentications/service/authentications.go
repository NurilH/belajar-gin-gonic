package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/NurilH/belajar-gin-gonic/model"
	"github.com/NurilH/belajar-gin-gonic/module/authentications"
	"github.com/NurilH/belajar-gin-gonic/module/users"
	"github.com/NurilH/belajar-gin-gonic/pkg/common/constants"
	"github.com/NurilH/belajar-gin-gonic/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type authenticationsService struct {
	redisAuth authentications.AuthRedisRepository
	repoAuth  authentications.AuthRepository
	repoUser  users.UsersRepository
}

func NewAuthService(
	redisAuth authentications.AuthRedisRepository, repoAuth authentications.AuthRepository, repoUser users.UsersRepository) authentications.AuthenticationsService {
	return &authenticationsService{
		redisAuth: redisAuth,
		repoAuth:  repoAuth,
		repoUser:  repoUser,
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
	res, redisErr := a.redisAuth.GetKey(c, constants.KeyRedisLogin+req.Email)
	if redisErr != nil || res == nil {
		user, err := a.repoUser.GetUserByEmail(c, req.Email)
		if err != nil {
			return result, err
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
		if err != nil {
			return result, err
		}

		// set time out JWT
		timeout := config.EnvAsDuration("JWT_TIMEOUT", 30*time.Minute)
		exp := time.Now().Add(timeout).Unix()

		// Create the Claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": user.ID,
			"email":   user.Email,
			"exp":     exp,
		})

		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString([]byte(config.Env("SECRET_KEY")))
		if err != nil {
			return result, err
		}

		redisExp := config.EnvAsDuration("REDIS_AUHT_EXP", 10*time.Minute)
		errRedis := a.redisAuth.Save(c, constants.KeyRedisLogin+req.Email, tokenString, redisExp)
		if errRedis != nil {
			return result, fmt.Errorf("invalid save redis %w", errRedis)
		}

		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("Authorization", tokenString, int(timeout.Seconds()), "", "", false, true)

		result.Token = &tokenString
		return result, err
	}
	result = *res
	return
}
