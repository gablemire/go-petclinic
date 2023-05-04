package config

type AppConfig struct {
	HttpServerPort int

	MongoHost     string
	MongoPort     int
	MongoUsername string
	MongoPassword string
}

var appConfig *AppConfig

func BootstrapConfig() *AppConfig {
	appConfig = &AppConfig{
		HttpServerPort: 8080,

		MongoHost:     "localhost",
		MongoPort:     27017,
		MongoUsername: "root",
		MongoPassword: "example",
	}

	return appConfig
}

func GetAppConfig() *AppConfig {
	if appConfig == nil {
		panic("AppConfig not bootstrapped")
	}

	return appConfig
}
