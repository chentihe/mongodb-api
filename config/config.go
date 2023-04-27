package config

import "github.com/spf13/viper"

type Config struct {
	DataBase DataBase `mapstructure:"database"`
}

type DataBase struct {
	Name     string `mapstructure:"name"`
	UserName string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("yaml")
	viper.SetConfigName("env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
