package config

import (
	"github.com/redis/go-redis/v9"
)

func GetRDB() *redis.Client {
	return rdb
}

func GetAppPort() string {
	return appPort
}

func GetUrlAccountService() string {
	return urlAccountService
}
