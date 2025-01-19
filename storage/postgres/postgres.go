package postgres

import (
	"database/sql"
	"fmt"
	"microsvc/api-gateway/data"
	"microsvc/common/utils"
	"os"

	_ "github.com/lib/pq"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	db *sql.DB
	// dsn    string
	logger *utils.CustomLogger
	config Config
}

type Config struct {
	dbName string
	port   string
	host   string
	user   string
	pswd   string
}

func NewStorage(logger *utils.CustomLogger, config Config) *Storage {
	return &Storage{
		logger: logger,
		config: config,
	}
}

func FormConfig() Config {
	return Config{
		dbName: os.Getenv("POSTGRES_NAME"),
		port:   os.Getenv("POSTGRES_PORT"),
		host:   os.Getenv("POSGRES_HOST"),
		user:   os.Getenv("POSTGRES_USER"),
		pswd:   os.Getenv("POSTGRES_PSWD"),
	}
}

func (ps *Storage) Migrate() {
	var err error

	ps.logger.Info("Start migration")

	if err = ps.createDBIfNotExist(); err != nil {
		ps.logger.Fatal("Can't create DB: %v", err)
	}

	ps.logger.Info("Conn to target DB: %s", ps.config.dbName)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", ps.config.host, ps.config.user, ps.config.pswd, ps.config.dbName, ps.config.port)

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		ps.logger.Fatal("Can't conn to DB: %v", err)
	}
	sqlDB, _ := gormDB.DB()
	defer sqlDB.Close()

	ps.logger.Info("Migrate...")

	if err = gormDB.AutoMigrate(&data.User{}); err != nil {
		ps.logger.Fatal("Migration error: %s", err)
	}

	ps.logger.Info("Migration succesful")
}

func (ps *Storage) createDBIfNotExist() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=%s sslmode=disable", ps.config.host, ps.config.user, ps.config.pswd, ps.config.port)

	ps.logger.Info("Conn to postgres DB")

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		ps.logger.Error("Can't conn to postgres DB: %v", err)
		return err
	}
	defer db.Close()

	ps.logger.Info("Check existing needed DB %s", ps.config.dbName)

	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)", ps.config.dbName).Scan(&exists)
	if err != nil {
		ps.logger.Error("Can't check existing needed DB: %v", err)
		return err
	}

	if !exists {
		ps.logger.Info("Target DB %s not found, creating...", ps.config.dbName)

		_, err = db.Exec("CREATE DATABASE " + ps.config.dbName)
		if err != nil {
			ps.logger.Error("Can't create DB %s: %v", ps.config.dbName, err)
			return err
		}

		ps.logger.Info("%s succesfuly created", ps.config.dbName)
	}

	return nil
}

func (ps *Storage) ConnToDB() error {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", ps.config.host, ps.config.user, ps.config.pswd, ps.config.dbName, ps.config.port)

	ps.db, err = sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	return nil
}

func (ps *Storage) Close() error {
	return ps.db.Close()
}

func (ps *Storage) Ping() error {
	return ps.db.Ping()
}
