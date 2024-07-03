package config

import (
	"cmp"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Host           string
	Port           string
	AllowedOrigins []string
	Local          bool
	DSN            string
}

func New() Config {
	host := getEnvDefault("HOST", "0.0.0.0")
	port := getEnvDefault("PORT", "8080")
	return Config{
		Host: host,
		Port: port,
		AllowedOrigins: strings.Split(
			getEnvDefault("ALLOWED_ORIGINS", fmt.Sprintf("http://%s:%s,https://%s:%s", host, port, host, port)), ",",
		),
		Local: getEnvDefault("LOCAL", "false") == "true",
		DSN:   getEnvDefault("DSN", "postgresql://user:password@localhost/data?sslmode=disable"),
	}
}

func getEnvDefault(key, def string) string {
	return cmp.Or(os.Getenv(key), def)
}
