package config

import (
	"os"

	"github.com/joho/godotenv"
)

const (
	env      = "ENVIRONMENT"
	name     = "APP_NAME"
	port     = "APP_PORT"
	dbUri    = "DB_URI"
	dbName   = "DB_NAME"
	redisUri = "REDIS_URI"
	redisKey = "REDIS_KEY"
	apiKey   = "API_KEY"
)

func Env() string      { return os.Getenv(env) }
func Name() string     { return os.Getenv(name) }
func Port() string     { return os.Getenv(port) }
func DbUri() string    { return os.Getenv(dbUri) }
func DbName() string   { return os.Getenv(dbName) }
func RedisUri() string { return os.Getenv(redisUri) }
func RedisKey() string { return os.Getenv(redisKey) }
func ApiKey() string   { return os.Getenv(apiKey) }

func init() {
	godotenv.Load()
}
