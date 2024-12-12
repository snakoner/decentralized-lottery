package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/sirupsen/logrus"
	"github.com/snakoner/lottery/internal/app"
	"github.com/snakoner/lottery/internal/config"
)

var (
	configPath = "./config.json"
)

func main() {
	logger := logrus.New()
	config, err := config.New(configPath)
	if err != nil {
		logger.Error(err)
	}

	logLevel, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		logLevel = logrus.DebugLevel
	}

	logger.SetLevel(logLevel)

	app, err := app.New(config, logger)
	if err != nil {
		logger.Error(err)
		return
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-c
		logger.Info("Got interrupt signal")
		cancel()
	}()

	if app.Run(ctx) != nil {
		return
	}
}
