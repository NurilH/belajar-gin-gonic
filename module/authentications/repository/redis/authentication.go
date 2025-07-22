package redis

import (
	"time"

	"github.com/NurilH/belajar-gin-gonic/model"
	"github.com/NurilH/belajar-gin-gonic/module/authentications"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type authRedisRepository struct {
	rc *redis.Client
}

func NewAuthRedisRepository(rc *redis.Client) authentications.AuthRedisRepository {
	return authRedisRepository{
		rc: rc,
	}
}

func (r authRedisRepository) Save(ctx *gin.Context, redisKey string, value string, exp time.Duration) error {
	err := r.rc.Set(ctx, redisKey, value, exp).Err()
	return err
}

func (r authRedisRepository) GetKey(ctx *gin.Context, redisKey string) (res *model.LoginRespons, err error) {
	val, err := r.rc.Get(ctx, redisKey).Result()
	if err != nil {
		return nil, err
	}
	// json.Unmarshal(val, res)
	res = new(model.LoginRespons)
	res.Token = &val
	return
}
