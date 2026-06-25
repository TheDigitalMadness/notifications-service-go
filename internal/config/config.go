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
		// TODO: заменить на корректное завершение. Например, через fx.Shutdowner.Shutdown
		panic(err)
	}

	return &cfg
}
