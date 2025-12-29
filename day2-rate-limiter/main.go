// Create rate limiter with N requests per second
// Simulate multiple requests
// Print allowed or rejected
package main

import (
	"fmt"
	"time"

	"github.com/Partho-Mal/go-core-systems/day2-rate-limiter/ratelimiter"
)

func main() {
	// allow 5 requests per second
	rl := ratelimiter.NewRateLimiter(5)

	clientID := "user-123"

	// simulate 10 requests
	for i := 1; i <= 10; i++ {
		allowed := rl.Allow(clientID)
		fmt.Println("request", i, "allowed:", allowed)
		time.Sleep(100 * time.Millisecond)
	}
}
