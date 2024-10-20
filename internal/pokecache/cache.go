package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	items map[string]cacheEntry
	mu    *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		items: make(map[string]cacheEntry),
		mu:    &sync.Mutex{},
	}

	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, exists := c.items[key]
	if !exists {
		entry := cacheEntry{createdAt: time.Now(), val: val}
		c.items[key] = entry
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, exists := c.items[key]
	if exists {
		return entry.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, value := range c.items {
		if value.createdAt.Before(now.Add(-last)) {
			delete(c.items, key)
		}
	}
}
