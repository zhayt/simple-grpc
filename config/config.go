package config

import (
	"fmt"
	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
	"log"
	"net"
)

type Config struct {
	Port         string `env:"PORT" envDefault:"5005"`
	DBConnection string `env:"DB_CONN"`
	Address      string
}

func NewConfig() (*Config, error) {
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("couldn't parse config: %w", err)
	}

	cfg.Address = net.JoinHostPort("", cfg.Port)
	return &cfg, nil
}

func MustPrepareEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
	}
}
