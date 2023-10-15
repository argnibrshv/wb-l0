package cache

import (
	"sync"
)

// кэш для хранения данных в оперативной памяти
type Cache struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewCache() *Cache {
	cache := new(Cache)
	cache.data = make(map[string]string)
	return cache
}

func (c *Cache) AddOrderToCache(id string, data string) {
	c.mu.Lock()
	c.data[id] = data
	c.mu.Unlock()
}

func (c *Cache) GetOrderByID(id string) (string, bool) {
	c.mu.RLock()
	data, ok := c.data[id]
	c.mu.RUnlock()
	if !ok {
		return data, false
	}
	return data, true
}
