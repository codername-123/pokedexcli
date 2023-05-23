package pokecache

import (
	"sync"
	"time"
)

// Cache -
type Cache struct {
	cache map[string]cacheEntry
	mux   *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// NewCache -
func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}

	go c.deleteLoop(interval)

	return c
}

// Add -
func (c *Cache) Add(key string, value []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	val, ok := c.cache[key]
	return val.val, ok
}

func (c *Cache) deleteLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.delete(time.Now().UTC(), interval)
	}
}

func (c *Cache) delete(now time.Time, interval time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	// t := time.Now().UTC().Add(-interval)
	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-interval)) {
			delete(c.cache, k)
		}
	}
}
