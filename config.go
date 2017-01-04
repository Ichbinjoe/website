package main

import (
	"github.com/spf13/viper"
)

func SetUpConfiguration() error {
	viper.AddConfigPath("/etc/ibjio")
	viper.AddConfigPath(".")
	return viper.ReadInConfig()
}
