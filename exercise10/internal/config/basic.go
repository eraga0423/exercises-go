package config

import "os"

type Config struct {
	//ENV   string
	API   APIConfig
	Redis RedisConfig
}

type APIConfig struct {
	Rest APIRestConfig
}
type APIRestConfig struct {
	Host string
	Port string
}
type RedisConfig struct {
	RedisName string
	Host      string
}

const NameRedis = "leaderBoard"

func NewConfig() *Config {
	return &Config{
		API: APIConfig{
			Rest: APIRestConfig{
				Host: "localhost",
				Port: os.Getenv("PORT"),
			},
		},
		Redis: RedisConfig{
			Host:      os.Getenv("REDIS_HOST"),
			RedisName: NameRedis,
		},
	}

}
