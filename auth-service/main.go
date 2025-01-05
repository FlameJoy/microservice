package main

import (
	"log" // Путь к сгенерированным файлам gRPC
	"microsvc/common/utils"
	"os"
)

func main() {
	// Логирование и старт сервера
	logger := utils.NewLogger(utils.INFO, log.New(os.Stdout, "auth-svc ", log.LstdFlags), false)

	// Запуск gRPC сервера на порту 50051
	StartGRPCServer(":50052", logger)
}
