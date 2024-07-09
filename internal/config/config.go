package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress  string
	WorkerPoolSize int
	QueueSize      int
}

func Load() (*Config, error) {
	viper.SetDefault("SERVER_ADDRESS", ":8080")
	viper.SetDefault("WORKER_POOL_SIZE", 10)
	viper.SetDefault("QUEUE_SIZE", 10000)

	viper.AutomaticEnv()

	return &Config{
		ServerAddress:  viper.GetString("SERVER_ADDRESS"),
		WorkerPoolSize: viper.GetInt("WORKER_POOL_SIZE"),
		QueueSize:      viper.GetInt("QUEUE_SIZE"),
	}, nil
}
