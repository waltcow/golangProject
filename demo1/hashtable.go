package demo1

import "errors"

var (
	ErrNotFound = errors.New("value is not found")
)

type  HashTable interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
}

