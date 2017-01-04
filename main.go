package main

import (
	reqman "./reqman"
	"github.com/spf13/viper"
	"github.com/uber-go/zap"
	"net/http"
	"time"
)

var logger zap.Logger

func main() {

	logger = CreatePreConfigLogger()

	logger.Info(":: Setting up configuration")
	err := SetUpConfiguration()
	if err != nil {
		logger.Panic("failed to read configuration", zap.Error(err))
	}

	logger.Info(":: Creating runtime logger")
	applogger, err := CreateRootLogger()
	if err != nil {
		logger.Panic("failed to create application logger", zap.Error(err))
	}

	logger.Info(":: Constructing server")

	handler := CreateRequestTimer

	handler := httpserver.CreatePipeline(httpserver.CreateRequestLogger(applogger), httpserver.CreateRequestTimer(), httpserver.WrapHttpHandler(http.FileServer(http.Dir("/home/joe"))))

	bindAddr := viper.GetString("bindAddr")

	server := http.Server{
		Addr:         bindAddr,
		Handler:      handler,
		ReadTimeout:  4 * time.Second,
		WriteTimeout: 4 * time.Second,
	}

	logger.Info(":: Starting hosting on " + bindAddr)

	err = server.ListenAndServe()

	panic(err)
}
