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
	mu    sync.Mutex
	items map[string]cacheEntry
}
