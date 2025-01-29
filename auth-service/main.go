package main

import (
	"log"
	"microsvc/common/utils"
	"microsvc/storage/postgres"

	"os"
)

func main() {
	logger := utils.NewLogger(utils.INFO, log.New(os.Stdout, "auth-svc ", log.LstdFlags), false)

	if err := utils.LoadEnv("../.env"); err != nil {
		logger.Fatal("LoadEnv error: %s", err)
	}

	storage := postgres.NewStorage(logger, postgres.FormConfig())
	sqlDir := "../storage/postgres/sql"

	// storage := clickhouse.NewStorage(logger, clickhouse.FormConfig())
	// sqlDir := "../storage/clickhouse/sql"

	if err := storage.Migrate(sqlDir); err != nil {
		logger.Fatal("DB error: %v", err)
	}

	if err := storage.ConnToDB(); err != nil {
		logger.Fatal("DB error: %v", err)
	}

	StartGRPCServer(":50052", logger, storage)
}
