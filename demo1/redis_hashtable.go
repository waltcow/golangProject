package demo1

import (
	"fmt"
	"gopkg.in/redis.v3"
)

type RedisClientWrapper struct {
	Client *redis.Client
}

func (r RedisClientWrapper) Get(key string) ([]byte, error) {
	val, ok := r.Client.Get(key).Result()
	if ok != nil {
		return nil, fmt.Errorf("Error: %s", ok)
	}
	return []byte(val), nil
}

func (r RedisClientWrapper) Set(key string, val []byte) error {
	err := r.Client.Set(key, val, 0).Err()
	return err
}

