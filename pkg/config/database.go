package config

import (
	"fmt"
	"time"

	"github.com/NurilH/belajar-gin-gonic/pkg/common/constants"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Host              string
	Port              int
	User              string
	Password          string
	DatabaseName      string
	MaxIdleConnection int
	MaxIdleTime       time.Duration
	MaxOpenConnection int
	MaxLifetime       time.Duration
}

func NewDBGormV2(c *Config) (*gorm.DB, error) {
	return newPostgresConnection(&c.MainDB)
}

func newPostgresConnection(c *Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"sslmode=disable host=%s port=%d user=%s password='%s' dbname=%s timezone=%s",
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.DatabaseName,
		constants.LocationAsiaJakarta,
	)

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(c.MaxIdleConnection)
	sqlDB.SetMaxOpenConns(c.MaxOpenConnection)
	sqlDB.SetConnMaxLifetime(c.MaxLifetime)

	return gormDB, nil
}
