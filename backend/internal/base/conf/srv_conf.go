package conf

var ServerConfigs = initServerConfig()

type ServerConfig struct {
	Addr string
}

func initServerConfig() ServerConfig {
	return ServerConfig{
		Addr: getEnv("SERVER_ADDR", ":8080"),
	}
}
