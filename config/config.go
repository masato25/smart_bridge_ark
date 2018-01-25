package config

import (
	"github.com/spf13/viper"
)

var config ViperConfig

func MyConfig() ViperConfig {
	return config
}

func ReadConf(confpath ...string) (err error) {
	viper.SetConfigName("conf")
	if len(confpath) != 0 {
		for _, cpath := range confpath {
			viper.AddConfigPath(cpath)
		}
	} else {
		// use default
		viper.AddConfigPath("./config")
		viper.AddConfigPath(".")
	}
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	config = config.SetAll()
	return
}
