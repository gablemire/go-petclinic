package config

type AppConfig struct {
	HttpServerPort int
}

var appConfig *AppConfig

func BootstrapConfig() *AppConfig {
	appConfig = &AppConfig{
		HttpServerPort: 8080,
	}

	return appConfig
}

func GetAppConfig() *AppConfig {
	if appConfig == nil {
		panic("AppConfig not bootstrapped")
	}

	return appConfig
}
