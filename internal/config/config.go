package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DB    DB
	HTTP  HTTP
	Kafka Kafka
}

func MustMakeConfig() *Config {
	cfg := Config{}

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}
