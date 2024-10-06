package configs

type StorageConfig struct {
	RedisHost     string `env:"REDIS_HOST,required"`
	RedisPort     string `env:"REDIS_PORT,required"`
	RedisPassword string `env:"REDIS_PASSWORD,required"`
	RedisDB       int    `env:"REDIS_DB_NUMBER,required"`
}
