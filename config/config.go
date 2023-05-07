package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Port     string   `mapstructure:"port"`
	DataBase DataBase `mapstructure:"database"`
	Jwt      Jwt      `mapstructure:"jwt"`
}

type DataBase struct {
	Name     string `mapstructure:"name"`
	UserName string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Cluster  string `mapstructure:"cluster"`
}

type Jwt struct {
	PrivateKey string        `mapstructure:"private_key"`
	PublicKey  string        `mapstructure:"public_key"`
	ExpiresIn  time.Duration `mapstructure:"expired_in"`
	MaxAge     int           `mapstructure:"max_age"`
}

func LoadConfig(paths ...string) (config *Config, err error) {
	for _, path := range paths {
		viper.AddConfigPath(path)
	}
	viper.SetConfigType("yaml")
	viper.SetConfigName("env")

	viper.SetDefault("port", "8080")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err = viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return config, nil
}
