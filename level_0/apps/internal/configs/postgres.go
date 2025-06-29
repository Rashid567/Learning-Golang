package configs

type PostgresConfig struct {
	Host         string `env:"HOST,required"`
	Port         int    `env:"PORT,required"`
	Login        string `env:"LOGIN,required"`
	Pasword      string `env:"PASSWORD,required"`
	DatabaseName string `env:"DATABASE_NAME,required"`
	SSLMode      string `env:"SSL_MODE,required"`
}
