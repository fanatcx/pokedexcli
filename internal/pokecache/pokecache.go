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
		cache: make(map[string]cacheEntry, 0),
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
	_, exist := c.cache[key]
	if exist {
		return c.cache[key].val, true
	}
	return nil, false
}


func (c *Cache) reapLoop(interval time.Duration) {
	//ch := make(chan)
	ch = time.Ticker(time.Duration)
	//????
	for key, _ := range c.cache {
		
		if interval
		if time.Time.Compare(nan, duration) {

		}
	}

}

// Cache functions //////////////////////







