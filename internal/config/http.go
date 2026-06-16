package config

type HTTP struct {
	Port int `env:"NODE_PORT" env-default:"3000"`
}
