package clickhouse

import (
	"context"
	"fmt"
	"log"
	"time"

	ch "github.com/ClickHouse/ch-go"
)

type Storage struct {
	db *ch.Client
}

func NewStorage() *Storage {
	return &Storage{}
}

func (clickhouse *Storage) ConnToDB() error {
	conn, err := ch.Dial(context.Background(), ch.Options{
		Address: "localhost:9000", // Адрес ClickHouse сервера
		// Укажите параметры пользователя, если они необходимы
		Database: "default",
		User:     "default",
		Password: "",
	})
	if err != nil {
		log.Fatalf("Не удалось подключиться к ClickHouse: %v", err)
		return err
	}
	// defer conn.Close()

	fmt.Println("Соединение установлено с clickhouse")

	clickhouse.db = conn
	return nil
}

// Close завершает соединение с ClickHouse
func (clickhouse *Storage) Close() error {
	if err := clickhouse.db.Close(); err != nil {
		log.Printf("Ошибка закрытия соединения: %v", err)
		return err
	}
	log.Println("Соединение с ClickHouse закрыто.")
	return nil
}

// Ping проверяет доступность ClickHouse
func (clickhouse *Storage) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := clickhouse.db.Ping(ctx)
	if err != nil {
		return fmt.Errorf("ClickHouse не отвечает: %w", err)
	}
	return nil
}
