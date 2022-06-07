package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`
	Port     string `mapstructure:"PORT"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config *Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	port := os.Getenv("PORT")
	if port != "" {
		config.Port = port
	}

	dbConfig := os.Getenv("DB_SOURCE")
	if dbConfig != "" {
		config.DBSource = dbConfig
	}

	return config, nil
}
