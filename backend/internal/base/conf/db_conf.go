package conf

import "fmt"

var DBConfigs = initDBConfig()

type dbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Address  string
	DBName   string
}

func initDBConfig() *dbConfig {
	return &dbConfig{
		Host:     getEnv("PUBLIC_HOST", "http://localhost"),
		Port:     getEnv("DB_PORT", "3306"),
		User:     getEnv("DB_USER", "root"),
		Password: getEnv("DB_PASSWORD", "root"),
		Address:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:   getEnv("DB_NAME", "social"),
	}
}
