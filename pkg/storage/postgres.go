package storage

import (
	"Practice.Golang/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDb() (*gorm.DB, error) {
	cfgs, err := config.NewConfigs()

	if err != nil {
		panic(err)
	}

	return gorm.Open(postgres.Open(cfgs.Database.DbConnectionString), &gorm.Config{})
}
