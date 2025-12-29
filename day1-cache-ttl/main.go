// Create cache
// Set value with TTL
// Get value immediately
// Wait for expiration
// Try Get again
package main

import (
	"fmt"
	"time"

	cache "github.com/Partho-Mal/go-core-systems/day1-cache-ttl/cache"
)

func main() {
	c := cache.NewCache()

	c.Set("name", "partho", 2*time.Second)

	val, ok := c.Get("name")
	fmt.Println(val, ok)

	time.Sleep(3 * time.Second)

	val, ok = c.Get("name")
	fmt.Println(val, ok)
}
