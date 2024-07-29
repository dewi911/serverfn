package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	ServerAddress  string
	WorkerPoolSize int
	QueueSize      int
}

type Config1 struct {
	DB     Postgres
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	WorkerPoolSize int
	QueueSize      int
}

type Postgres struct {
	Host     string
	Port     int
	Username string
	Name     string
	SSLMode  string
	Password string
}

func New(folder, filename string) (*Config1, error) {
	cfg := new(Config1)

	viper.AddConfigPath(folder)
	viper.SetConfigName(filename)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	if err := envconfig.Process("DB", &cfg.DB); err != nil {
		log.Printf("Error processing DB_HOST: %v", err)
		return nil, err
	}

	return cfg, nil
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
