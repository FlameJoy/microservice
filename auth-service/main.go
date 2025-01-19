package main

import (
	"log" // Путь к сгенерированным файлам gRPC
	"microsvc/common/utils"
	"microsvc/storage/postgres"
	"os"
)

func main() {
	// Логирование и старт сервера
	logger := utils.NewLogger(utils.INFO, log.New(os.Stdout, "auth-svc ", log.LstdFlags), false)

	if err := utils.LoadEnv("../.env"); err != nil {
		logger.Fatal("LoadEnv error: %s", err)
	}

	// DB
	config := postgres.FormConfig()
	storage := postgres.NewStorage(logger, config)
	storage.Migrate()
	if err := storage.ConnToDB(); err != nil {
		logger.Fatal("DB error: %v", err)
	}

	// Запуск gRPC сервера на порту 50051
	StartGRPCServer(":50052", logger, storage)
}
