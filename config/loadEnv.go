package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return err
	}

	appPort = os.Getenv(APP_PORT)
	urlRedis = os.Getenv(URL_REDIS)
	urlAccountService = os.Getenv(URL_ACCOUNT_SERVICE)
	urlShopService = os.Getenv(URL_SHOP_SERVICE)
	urlProductService = os.Getenv(URL_PRODUCT_SERVICE)
	urlFileService = os.Getenv(URL_FILE_SERVICE)
	urlSearchService = os.Getenv(URL_SEARCH_SERVICE)
	urlPayment = os.Getenv(URL_PAYMENT_SERVICE)
	urlOrderService = os.Getenv(URL_ORDER_SERVICE)

	return nil
}
