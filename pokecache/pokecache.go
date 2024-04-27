package pokecache

import (
	"time"
)

type Cache struct {
	Cachebody map[string]cacheEntry
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		Cachebody: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.Cachebody[key] = cacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) (val []byte, b bool) {

	buff := c.Cachebody[key].val
	if buff != nil {
		return buff, true
	} else {
		return nil, false
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval) // will run every "interva"
	}

}

func (c *Cache) reap(interval time.Duration) {
	Timepast := time.Now().UTC().Add(-interval)
	for k, v := range c.Cachebody {
		if v.createdAt.Before(Timepast) {
			delete(c.Cachebody, k)
		}
	}

}
