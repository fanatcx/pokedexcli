package pokecache
import (
	"time"
	"sync"
)

// add mutex
type Cache struct {
	cache       map[string]cacheEntry
	mu          sync.Mutex
	
}

type cacheEntry struct {
	createdAt 	time.Time 
	val 		[]byte 		

}

func NewCache(interval time.Duration) *Cache {
	newCache := Cache{
		cache: make(map[string]cacheEntry),
	}

	go newCache.reapLoop(interval)
	return &newCache
}

// Cache functions //////////////////////

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	cacheEntry := cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.cache[key] = cacheEntry
}


func (c *Cache) Get(key string) ([]byte, bool) {
	// I thought reading from a map is safe?
	c.mu.Lock()
	defer c.mu.Unlock()

	cacheEntry, exist := c.cache[key]
	if exist {
		return cacheEntry.val, true
	}

	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// Go sends the current time (time.Now()) automatically to that ticker. Didnt know this
	for range ticker.C {
		c.mu.Lock()
		
		for key, entry := range c.cache {
			age := entry.createdAt
			if time.Since(age) > interval {
				delete(c.cache, key)
			}
			
		}
		// done scanning
		c.mu.Unlock()
	}

}

// Cache functions //////////////////////







