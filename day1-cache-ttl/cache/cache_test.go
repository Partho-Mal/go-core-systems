package cache_test

import (
	"testing"
	"time"

	"github.com/Partho-Mal/go-core-systems/day1-cache-ttl/cache"
)

func TestSetAndGet(t *testing.T) {
	c := cache.NewCache()

	c.Set("name", "partho", 2*time.Second)

	val, ok := c.Get("name")
	if !ok {
		t.Fatalf("expected key to exist")
	}

	if val != "partho" {
		t.Fatalf("expected partho, got %v", val)
	}
}

func TestExpiration(t *testing.T) {
	c := cache.NewCache()

	c.Set("temp", "data", 500*time.Millisecond)

	time.Sleep(1 * time.Second)

	_, ok := c.Get("temp")
	if ok {
		t.Fatalf("expected key to expire")
	}
}

