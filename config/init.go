package config

import (
	"flag"
)

func init() {
	var migrate bool = false
	flag.BoolVar(&migrate, "db", true, "Migrate Database?")

	loadEnv()
	connectRedis()
}
