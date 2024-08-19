package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	Env  string `yaml:"env" env:"ENV" env-default:"dev"`
	Http `yaml:"http"`
	PgDb `yaml:"apigatway"`
}

type PgDb struct {
	PgHost     string `yaml:"host" env:"HOST" env-default:"localhost"`
	PgPort     string `yaml:"pg_port" env:"PORT" env-default:"5432"`
	PgName     string `yaml:"dbname" env:"NAME" env-default:"apigatway"`
	PgUser     string `yaml:"username" env:"USER" env-default:"alex"`
	PgPassword string `yaml:"password" env:"PASSWORD" env-default:"secret"`
	SSLMode    string `yaml:"sslmode" env:"SSL_MODE" env-default:"disable"`
}

type Http struct {
	Port        string `yaml:"port" env:"PORT" env-default:"8080"`
	Timeout     string `yaml:"timeout" env:"TIMEOUT" env-default:"4s"`
	IdleTimeout string `yaml:"idle_timeout" env:"IDLE_TIMEOUT" env-default:"10s"`
}

func MustLoad() Config {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		log.Fatalf("CONFIG_PATH not set %s", configPath)
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist: %s", err)
	}
	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("Failed to load config: %s", err)
	}
	return cfg
}
