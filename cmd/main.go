package main

import (
	"os"
	"strings"
	"sync"

	"github.com/jmvdr-iscte/ImageSearch/config"
	"github.com/jmvdr-iscte/ImageSearch/internal/api"
	"github.com/jmvdr-iscte/ImageSearch/pkg/logger"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"err": err,
		}).Error("Can't load config from .env. Problem with .env, or the server is in production environment.")
		return
	}

	config := config.ApiEnvConfig{
		Env:     strings.ToUpper(os.Getenv("ENV")),
		Port:    os.Getenv("PORT"),
		Version: os.Getenv("VERSION"),
	}

	logger.Log.WithFields(logrus.Fields{
		"env":     config.Env,
		"version": config.Version,
		"port":    config.Port,
	}).Info("Loaded app config")

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		server := api.AppServer{}
		defer func() {
			if r := recover(); r != nil {
				server.OnShutdown()
			}
		}()

		server.Run(config)
	}()
	wg.Wait()

}
