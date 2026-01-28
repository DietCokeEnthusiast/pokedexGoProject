package pokecache

import (
	"sync"
	"time"

)

// start of cache that will be imported into our pokedex
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// mutex to protect. Read&write at same time can cause crashing in program
type Cache struct {
	mu 			sync.Mutex
	items 		map[string]cacheEntry
	interval 	time.Duration
}


func NewCache(interval time.Duration) *Cache {
	

	cache := &Cache{items: make(map[string]cacheEntry),
					interval: interval}

		go cache.reapLoop()


	return cache
}

func (cache *Cache) Add(key string,val []byte){
  cache.mu.Lock()

  defer cache.mu.Unlock()


	cache.items[key] = cacheEntry{createdAt: time.Now(), val: val}
	

}

func (cache *Cache) Get(key string) ([]byte, bool){
	cache.mu.Lock()

	defer cache.mu.Unlock()
  

	entry, found := cache.items[key]

	if !found {
		return nil, false
	}

  return entry.val, true

}

func (cache *Cache) reapLoop() {

	cache.mu.Lock()

	defer cache.mu.Unlock()

	for key, entry := range cache.items{
	if time.Since(entry.createdAt)>cache.interval{
		delete(cache.items, key)
	}
}

}
