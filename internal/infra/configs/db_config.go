package configs

type DBConfig struct {
	// Database
	DBHost     string `env:"DB_HOST,required"`
	DBPort     string `env:"DB_PORT,required"`
	DBDatabase string `env:"DB_DATABASE,required"`
	DBUsername string `env:"DB_USERNAME,required"`
	DBPassword string `env:"DB_PASSWORD,required"`
}
