package storage

import "time"

type Service interface {
	GetByKey(key string) ([]byte, error)
	Set(key string, value []byte, exp time.Duration) error
	GetKeys() []string
	RemoveKey(key string) error

	Close() error
}
