// Define Cache
//   Map of string to Item
//   Read write mutex

// Create NewCache
//   Initialize map
//   Start background cleaner
//   Return cache

// Set(key, value, ttl)
//   Lock for writing
//   Calculate expiration time
//   Store item in map
//   Unlock

// Get(key)
//   Lock for reading
//   Check if key exists
//   Unlock read lock

//   If not found
//     Return false

//   If expired
//     Lock for writing
//     Delete key
//     Unlock
//     Return false

// Return value and true


package cache

import (
	"sync"
	"time"
)

// Cache represents the in memory cache
type Cache struct {
	// map storing key to item
	items map[string]Item

	// mutex to protect map access
	mu sync.RWMutex
}

// NewCache creates and initializes the cache
func NewCache() *Cache {
	// initialize cache struct
	c := &Cache{
		items: make(map[string]Item),
	}

	// start background cleaner goroutine
	go c.startCleaner()

	return c
}

// Set inserts or updates a key with ttl
func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	// acquire write lock
	c.mu.Lock()
	defer c.mu.Unlock()

	// calculate expiration time
	expiration := time.Now().Add(ttl)

	// store item in map
	c.items[key] = Item{
		value:      value,
		expiration: expiration,
	}
}

// Get retrieves a value from cache
func (c *Cache) Get(key string) (interface{}, bool) {
	// acquire read lock
	c.mu.RLock()
	item, found := c.items[key]
	c.mu.RUnlock()

	// if key not found
	if !found {
		return nil, false
	}

	// check if item is expired
	if time.Now().After(item.expiration) {
		// acquire write lock to delete expired item
		c.mu.Lock()
		delete(c.items, key)
		c.mu.Unlock()

		return nil, false
	}

	// return value if valid
	return item.value, true
}
