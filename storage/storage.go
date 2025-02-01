package storage

import "database/sql"

type Storage interface {
	Close() error
	Ping() error
	ExecuteQuery(query string, args ...interface{}) (int64, error)
	ExecuteUpdate(table string, id int64, updates map[string]interface{}) error
	GetData(query string, args ...interface{}) ([]map[string]interface{}, error)
	DB() *sql.DB
}
