package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Environment string

	UserServiceHost string
	UserServicePort int

	PostServiceHost string
	PostServicePort int

	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string

	RedisHost string
	RedisPort int

	CtxTimeout int

	LogLevel  string
	HTTPPort  string
	HTTPHost  string
	SigninKey string
}

func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	c.UserServiceHost = cast.ToString(getOrReturnDefault("USERSERVICE_HOST", "localhost"))
	c.UserServicePort = cast.ToInt(getOrReturnDefault("USERSERVICE_PORT", ":9000"))
	c.PostServiceHost = cast.ToString(getOrReturnDefault("POSTSERVICE_HOST", "localhost"))
	c.HTTPHost = cast.ToString(getOrReturnDefault("HTTP_HOST", "localhost"))
	c.PostServicePort = cast.ToInt(getOrReturnDefault("POSTSERVICE_PORT", ":8000"))
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8080"))
	c.RedisHost = cast.ToString(getOrReturnDefault("REDIS_HOST", "localhost"))
	c.RedisPort = cast.ToInt(getOrReturnDefault("REDIS_PORT", "6379"))
	c.SigninKey = cast.ToString(getOrReturnDefault("SIGNIN_KEY", "sijxoxyffnfxemfhuoehmniihgs"))
	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres"))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "userss"))

	c.CtxTimeout = cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7))
	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
