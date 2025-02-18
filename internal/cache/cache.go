package cache

import (
	"time"
)

type CacheInterface interface {
	Get(key string) (string, error)
	Set(key string, value string, expiration time.Duration) error
	Delete(key string) error
	Close() error
}
