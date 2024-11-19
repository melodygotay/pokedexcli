package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu    sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(t time.Duration) *Cache {
	c := &Cache{
		cache: make(map[string]cacheEntry),
		mu:    sync.Mutex{},
	}
	c.reapLoop(t)
	return c
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	elem, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	return elem.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			c.mu.Lock()
			for key, entry := range c.cache {
				if entry.createdAt.Before(time.Now().Add(-interval)) {
					delete(c.cache, key)
				}
			}
			c.mu.Unlock()
		}
	}()
}
