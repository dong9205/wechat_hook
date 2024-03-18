package config

import (
	"github.com/spf13/viper"
)

func ConfigInit(config *string) error {
	viper.SetConfigFile(*config)
	return viper.ReadInConfig()
}
