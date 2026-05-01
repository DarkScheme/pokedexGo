package pokecache

import (
	"time"
	"sync"
	// "fmt"
)

type Cache struct {
	entries map[string]cacheEntry
	mut sync.Mutex 
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) *Cache {
	m := make(map[string]cacheEntry)
	newC := Cache{
		entries: m,
		mut: sync.Mutex{},
	}
	go newC.reapLoop(interval)
	// fmt.Printf("NewCache returning %p\n", &newC)
	return &newC
}

func (c *Cache) Add(key string, val []byte) {
	c.mut.Lock()
	defer c.mut.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	// fmt.Printf("Get called, c = %p\n", c)
	c.mut.Lock()
	defer c.mut.Unlock()

	val, ok := c.entries[key]
	if ok {
		return val.val, true
	} 
	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mut.Lock()
		for key, entry := range c.entries {
			if time.Since(entry.createdAt) > interval {
				delete(c.entries, key)
			}
		}
		c.mut.Unlock()
	}
}