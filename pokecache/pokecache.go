package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Cachebody map[string]cacheEntry
	Check     sync.Mutex // mutex
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache() Cache {
	return Cache{
		Cachebody: make(map[string]cacheEntry),
	}
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
