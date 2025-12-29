// Define RateLimiter
//   map of clientID to Bucket
//   mutex

// Allow(clientID)
//   get or create bucket
//   try to take token
//   if token available
//     return true
//   else
//     return false

package ratelimiter

import (
	"sync"
	"time"
)

// RateLimiter manages buckets for multiple clients
type RateLimiter struct {
	// map of clientID to token bucket
	buckets map[string]*Bucket

	// mutex to protect bucket map
	mu sync.Mutex

	// bucket capacity
	capacity int

	// refill interval
	refillInterval time.Duration
}

// NewRateLimiter initializes rate limiter
func NewRateLimiter(capacity int) *RateLimiter {
	return &RateLimiter{
		buckets:        make(map[string]*Bucket),
		capacity:       capacity,
		refillInterval: time.Second,
	}
}

// Allow checks if a request is allowed for a client
func (r *RateLimiter) Allow(clientID string) bool {
	r.mu.Lock()

	// get or create bucket for client
	bucket, exists := r.buckets[clientID]
	if !exists {
		bucket = NewBucket(r.capacity, r.refillInterval)
		r.buckets[clientID] = bucket
	}

	r.mu.Unlock()

	// non blocking token check
	return bucket.Allow()
}
