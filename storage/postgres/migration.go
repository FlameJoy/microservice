package postgres

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
)

func (ps *Storage) Migrate(migrationsDir string) error {
	var err error

	ps.logger.Info("Start migration")

	if err = ps.createDBIfNotExist(); err != nil {
		return err
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", ps.config.host, ps.config.user, ps.config.pswd, ps.config.dbName, ps.config.port)

	ps.db, err = sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	defer ps.db.Close()

	ps.logger.Info("Migrate...")

	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue // Пропускаем вложенные директории
		}

		ps.logger.Info("Execute %s...", file.Name())
		// Полный путь к SQL-файлу
		filePath := filepath.Join(migrationsDir, file.Name())
		if err := execSQLFile(ps.db, filePath); err != nil {
			return err
		}
	}

	ps.logger.Info("Migration succesful")

	return nil
}

func (ps *Storage) createDBIfNotExist() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=%s sslmode=disable", ps.config.host, ps.config.user, ps.config.pswd, ps.config.port)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)", ps.config.dbName).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		ps.logger.Info("Target DB %s not found, creating...", ps.config.dbName)

		_, err = db.Exec("CREATE DATABASE " + ps.config.dbName)
		if err != nil {
			return err
		}

		ps.logger.Info("%s DB succesfuly created", ps.config.dbName)
	}

	return nil
}

func execSQLFile(db *sql.DB, filepath string) error {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(content))
	if err != nil {
		return err
	}

	return nil
}
