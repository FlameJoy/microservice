package storage

type Storage interface {
	Close() error
	Ping() error
}
