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

func GetUrlShopService() string {
	return urlShopService
}

func GetUrlProductService() string {
	return urlProductService
}

func GetUrlOrderService() string {
	return urlOrderService
}

func GetUrlFileService() string {
	return urlFileService
}

func GetUrlSearchService() string {
	return urlSearchService
}

func GetUrlPayment() string {
	return urlPayment
}
