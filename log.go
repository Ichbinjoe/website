package main

import (
	"github.com/spf13/viper"
	"github.com/uber-go/zap"
	"os"
)

var logEncoder = zap.NewJSONEncoder(zap.RFC3339Formatter("timestamp"),
	zap.MessageKey("msg"),
	zap.LevelString("level"))

func CreatePreConfigLogger() zap.Logger {
	return zap.New(logEncoder)
}

func CreateRootLogger() (zap.Logger, error) {
	isProd := viper.GetBool("production")

	var logger zap.Logger
	if isProd {
		// elasticsearch?
		logger = zap.New(logEncoder, zap.AddStacks(zap.PanicLevel), zap.Output(os.Stderr))
	} else {
		logger = zap.New(logEncoder, zap.Development(), zap.AddStacks(zap.ErrorLevel), zap.Output(os.Stdout))
	}
	return logger, nil
}
