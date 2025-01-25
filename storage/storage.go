package storage

type Storage interface {
	Close() error
	Ping() error
	ExecuteQuery(query string, args ...interface{}) error
	GetData(query string, args ...interface{}) ([]map[string]interface{}, error)
}
