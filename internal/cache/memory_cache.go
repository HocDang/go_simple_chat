package cache

import (
	"sync"
	"time"
)

type memoryItem struct {
	value      string
	expiration int64
}

type MemoryCache struct {
	items map[string]memoryItem
	mu    sync.RWMutex
}

func InitMemoryCache() *MemoryCache {
	return &MemoryCache{
		items: make(map[string]memoryItem),
	}
}

func (m *MemoryCache) Get(key string) (string, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	item, exists := m.items[key]
	if !exists || (item.expiration > 0 && item.expiration < time.Now().Unix()) {
		return "", nil
	}
	return item.value, nil
}

func (m *MemoryCache) Set(key string, value string, expiration time.Duration) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	var exp int64
	if expiration > 0 {
		exp = time.Now().Add(expiration).Unix()
	}
	m.items[key] = memoryItem{
		value:      value,
		expiration: exp,
	}
	return nil
}

func (m *MemoryCache) Delete(key string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.items, key)
	return nil
}

func (m *MemoryCache) Close() error {
	return nil
}
