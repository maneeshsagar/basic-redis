// config/config.go
package config

import "os"

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

var RedisSettings = RedisConfig{
	Addr:     getEnv("REDIS_ADDR", "localhost:6379"),
	Password: "",
	DB:       0,
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
