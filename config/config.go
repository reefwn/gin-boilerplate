package config

import (
	"gin-boilerplate/helpers/logger"

	"github.com/spf13/viper"
)

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

func SetupConfig() error {
	var configuration *Configuration

	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		logger.Errorf("[SetupConfig]: ReadInConfig, %s", err)
		return err
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		logger.Errorf("[SetupConfig]: Unmarshal, %v", err)
		return err
	}

	return nil
}
