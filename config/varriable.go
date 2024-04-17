package config

import (
	"github.com/redis/go-redis/v9"
)

const (
	APP_PORT            = "APP_PORT"
	URL_REDIS           = "URL_REDIS"
	URL_ACCOUNT_SERVICE = "URL_ACCOUNT_SERVICE"
	URL_SHOP_SERVICE    = "URL_SHOP_SERVICE"
	URL_PRODUCT_SERVICE = "URL_PRODUCT_SERVICE"
	URL_ORDER_SERVICE   = "URL_ORDER_SERVICE"
)

var (
	urlAccountService string
	urlShopService    string
	urlProductService string
	urlOrderService   string

	appPort  string
	urlRedis string
	rdb      *redis.Client
)
