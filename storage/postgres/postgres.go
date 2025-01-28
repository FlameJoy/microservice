package postgres

import (
	"database/sql"
	"fmt"
	"microsvc/common/utils"
	"os"

	_ "github.com/lib/pq"
)

type Storage struct {
	db     *sql.DB
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
		dbName: os.Getenv("PS_NAME"),
		port:   os.Getenv("PS_PORT"),
		host:   os.Getenv("PS_HOST"),
		user:   os.Getenv("PS_USER"),
		pswd:   os.Getenv("PS_PSWD"),
	}
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

// ExecuteQuery - INSERT, UPDATE, DELETE
func (ps *Storage) ExecuteQuery(query string, args ...interface{}) error {
	_, err := ps.db.Exec(query, args...)
	return err
}

// GetData - SELECT
func (ps *Storage) GetData(query string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := ps.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}

	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))

		for i := range values {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, err
		}

		row := make(map[string]interface{})
		for i, col := range columns {
			val := values[i]
			if b, ok := val.([]byte); ok {
				row[col] = string(b)
			} else {
				row[col] = val
			}
		}

		result = append(result, row)
	}

	return result, nil
}
