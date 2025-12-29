// Define Bucket
//   capacity
//   current tokens
//   mutex
//   refill channel or ticker

// Create new bucket
//   set capacity
//   start refill goroutine

// Refill loop
//
//	every tick
//	  if tokens < capacity
//	    add token

package ratelimiter

import (
	"sync"
	"time"
)

// Bucket represents a token bucket for one client
type Bucket struct {
	// maximum tokens allowed
	capacity int

	// current available tokens
	tokens int

	// mutex to protect token count
	mu sync.Mutex

	// ticker for refilling tokens
	ticker *time.Ticker
}

// NewBucket creates a bucket and starts refill goroutine
func NewBucket(capacity int, refillInterval time.Duration) *Bucket {
	b := &Bucket{
		capacity: capacity,
		tokens:   capacity,
		ticker:   time.NewTicker(refillInterval),
	}

	// start background refill
	go b.refill()

	return b
}

// refill adds tokens periodically
func (b *Bucket) refill() {
	for range b.ticker.C {
		b.mu.Lock()

		// add token only if bucket is not full
		if b.tokens < b.capacity {
			b.tokens++
		}

		b.mu.Unlock()
	}
}

// Allow tries to consume one token
func (b *Bucket) Allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	// if no tokens left, reject request
	if b.tokens <= 0 {
		return false
	}

	// consume one token
	b.tokens--
	return true
}

