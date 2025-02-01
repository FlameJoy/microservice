package clickhouse

import (
	"database/sql"
	"fmt"
	"microsvc/common/utils"
	"os"

	_ "github.com/ClickHouse/clickhouse-go/v2"
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
		dbName: os.Getenv("CH_DBNAME"),
		port:   os.Getenv("CH_PORT"),
		host:   os.Getenv("CH_HOST"),
		user:   os.Getenv("CH_USER"),
		pswd:   os.Getenv("CH_PSWD"),
	}
}

func (ch *Storage) ConnToDB() error {
	var err error
	dsn := fmt.Sprintf("tcp://%s:%s?username=%s&password=%s&database=%s", ch.config.host, ch.config.port, ch.config.user, ch.config.pswd, ch.config.dbName)

	ch.db, err = sql.Open("clickhouse", dsn)
	if err != nil {
		return err
	}

	return nil
}

func (ch *Storage) Close() error {
	return ch.db.Close()
}

func (ch *Storage) Ping() error {

	return ch.db.Ping()
}

// ExecuteQuery - INSERT, UPDATE, DELETE
func (ch *Storage) ExecuteQuery(query string, args ...interface{}) error {
	_, err := ch.db.Exec(query, args...)
	return err
}

// GetData - SELECT
func (ch *Storage) GetData(query string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := ch.db.Query(query, args...)
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

func (ch *Storage) DB() *sql.DB {
	return ch.db
}
