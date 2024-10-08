package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	Env  string `yaml:"env" env:"ENV" env-default:"dev"`
	Http `yaml:"http"`
	PgDb `yaml:"pg-db"`
}

type PgDb struct {
	PgHost     string `yaml:"host" env:"PG_HOST" env-default:"postgresql"`
	PgPort     string `yaml:"pg_port" env:"PG_PORT" env-default:"5432"`
	PgName     string `yaml:"dbname" env:"PG_DB" env-default:"dbNamePg"`
	PgUser     string `yaml:"username" env:"PG_USER" env-default:"alex"`
	PgPassword string `yaml:"password" env:"PG_PASSWORD" env-default:"secret"`
	SSLMode    string `yaml:"sslmode" env:"PG_SSLMODE" env-default:"disable"`
}

type Http struct {
	Port        string `yaml:"port" env:"PORT" env-default:"8080"`
	Timeout     string `yaml:"timeout" env:"TIMEOUT" env-default:"4s"`
	IdleTimeout string `yaml:"idle_timeout" env:"IDLE_TIMEOUT" env-default:"10s"`
}

func MustLoad() Config {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		configPath = "/usr/local/src/local.yaml"
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
