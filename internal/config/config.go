package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	DB    DB
	HTTP  HTTP
	Kafka Kafka
}

func MustMakeConfig() *Config {
	_ = godotenv.Load()
	cfg := Config{}

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}
