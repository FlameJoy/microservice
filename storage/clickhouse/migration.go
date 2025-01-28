package clickhouse

import (
	"database/sql"
	"fmt"
	"microsvc/common/utils"
	"os"
	"path/filepath"
)

func (ch *Storage) Migrate(migrationsDir string) error {
	var err error

	ch.logger.Info("Start migration")

	if err = ch.createDBIfNotExist(); err != nil {
		return err
	}

	dsn := fmt.Sprintf("tcp://%s:%s?username=%s&password=%s", ch.config.host, ch.config.port, ch.config.user, ch.config.pswd)

	ch.db, err = sql.Open("clickhouse", dsn)
	if err != nil {
		return err
	}
	defer ch.db.Close()

	_, err = ch.db.Exec("USE " + ch.config.dbName)
	if err != nil {
		return err
	}

	ch.logger.Info("Migrate...")

	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue // Пропускаем вложенные директории
		}

		ch.logger.Info("Execute %s...", file.Name())
		// Полный путь к SQL-файлу
		filePath := filepath.Join(migrationsDir, file.Name())
		if err := utils.ExecSQLFile(ch.db, filePath); err != nil {
			return err
		}
	}

	ch.logger.Info("Migration succesful")

	return nil
}

func (ch *Storage) createDBIfNotExist() error {
	dsn := fmt.Sprintf("tcp://%s:%s?username=%s&password=%s", ch.config.host, ch.config.port, ch.config.user, ch.config.pswd)

	db, err := sql.Open("clickhouse", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	var exists bool
	err = db.QueryRow(`SELECT COUNT(*) > 0 FROM system.databases WHERE name = ?`, ch.config.dbName).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		ch.logger.Info("Target DB %s not found, creating...", ch.config.dbName)

		_, err = db.Exec("CREATE DATABASE " + ch.config.dbName)
		if err != nil {
			return err
		}

		ch.logger.Info("%s DB succesfuly created", ch.config.dbName)
	}

	return nil
}
