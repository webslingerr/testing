package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	// DebugMode indicates service mode is debug.
	DebugMode = "debug"
	// TestMode indicates service mode is test.
	TestMode = "test"
	// ReleaseMode indicates service mode is release.
	ReleaseMode = "release"
)

type Config struct {
	Environment string // debug, test, release

	ServerHost string
	ServerPort string

	PostgresHost     string
	PostgresUser     string
	PostgresDatabase string
	PostgresPassword string
	PostgresPort     string

	DefaultOffset int
	DefaultLimit  int

	PostgresMaxConnections int32
}

func Load() Config {

	if err := godotenv.Load("./app.env"); err != nil {
		fmt.Println("No .env file found")
	}

	cfg := Config{}

	cfg.ServerHost = cast.ToString(getOrReturnDefaultValue("SERVICE_HOST", "localhost"))
	cfg.ServerPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":4000"))

	cfg.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost"))
	cfg.PostgresPort = cast.ToString(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	cfg.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "shokhrukh"))
	cfg.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "12345"))
	cfg.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "client"))

	cfg.DefaultOffset = cast.ToInt(getOrReturnDefaultValue("OFFSET", 0))
	cfg.DefaultLimit = cast.ToInt(getOrReturnDefaultValue("LIMIT", 10))

	cfg.PostgresMaxConnections = 20

	return cfg
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
