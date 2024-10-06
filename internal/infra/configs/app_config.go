package configs

type AppConfig struct {
	// APP
	Environment string `env:"APP_ENV,required"`
	Name        string `env:"APP_NAME,required"`
	Port        string `env:"APP_PORT,required"`
	Version     string `env:"APP_VERSION,required"`
	BaseUrl     string `env:"APP_BASE_URL,required"`
}
