package config

type DB struct {
	Host     string `env:"DB_HOST" env-required:"true"`
	Port     uint   `env:"DB_PORT" env-required:"true"`
	User     string `env:"DB_USER" env-required:"true"`
	Password string `env:"DB_PASSWORD" env-required:"true"`
	DB       string `env:"DB_NAME" env-required:"true"`
}
