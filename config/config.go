package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const envPrefix = "BASSOON"

type Config struct {
	PortsFilepath string `envconfig:"PORTS_FILEPATH"   default:"ports.json"`
	HTTPPort      string `envconfig:"HTTP_PORT"        default:":8001"`
}

func Load() (*Config, error) {
	var cfg Config

	if err := godotenv.Load(); err != nil {
		log.Printf("WARN: failed to read env file: %s", err.Error())
	}

	err := envconfig.Process(envPrefix, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
