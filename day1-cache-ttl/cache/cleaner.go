// Start cleaner loop
//   Every fixed interval
//     Lock cache
//     For each item
//       If expired
//         Delete it
//     Unlock

package cache

import "time"

// startCleaner periodically removes expired keys
func (c *Cache) startCleaner() {
	// create ticker for periodic cleanup
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	// run forever
	for range ticker.C {
		// get current time
		now := time.Now()

		// acquire write lock
		c.mu.Lock()

		// iterate over all items
		for key, item := range c.items {
			// if item is expired
			if now.After(item.expiration) {
				// remove from cache
				delete(c.items, key)
			}
		}

		// release lock
		c.mu.Unlock()
	}
}
