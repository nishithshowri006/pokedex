package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache    map[string]cacheEntry
	mut      *sync.Mutex
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	var cache = make(map[string]cacheEntry)
	var mut = sync.Mutex{}
	c := &Cache{

		cache: cache,
		mut:   &mut,
	}
	go c.reapLoop(interval)
	return c
}
func (c *Cache) Add(key string, val []byte) {
	c.mut.Lock()
	defer c.mut.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mut.Lock()
	defer c.mut.Unlock()
	val, ok := c.cache[key]
	return val.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {

		c.mut.Lock()
		for k := range c.cache {
			if time.Now().Sub(c.cache[k].createdAt) >= c.interval {
				delete(c.cache, k)
			}
		}
		c.mut.Unlock()
	}

}
