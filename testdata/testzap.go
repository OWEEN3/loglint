package testdata

import (
	"go.uber.org/zap"
)

// Примеры взяты с задания
func TestLogsZap() {
	password := "secret123"
	apiKey := "apikey123"
	token := "token456"

	log, _ := zap.NewProduction()
	defer log.Sync()

	// Неправильные примеры логов
	log.Info("Starting server on port 8080")
	log.Error("Failed to connect to database")
	log.Info("запуск сервера")
	log.Error("ошибка подключения к базе данных")
	log.Info("server started!🚀")
	log.Error("connection failed!!!")
	log.Warn("warning: something went wrong...")
	log.Info("user password: " + password)
	log.Debug("api_key=" + apiKey)
	log.Info("token: " + token)

	// Правильные примеры логов
	log.Info("starting server on port 8080")
	log.Error("failed to connect to database")
	log.Info("starting server")
	log.Error("failed to connect to database")
	log.Info("server started")
	log.Error("connection failed")
	log.Warn("something went wrong")
	log.Info("user authenticated successfully")
	log.Debug("api request completed")
	log.Info("token validated")
}