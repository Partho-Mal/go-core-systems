package ratelimiter_test

import (
	"testing"
	"time"

	"github.com/Partho-Mal/go-core-systems/day2-rate-limiter/ratelimiter"
)

func TestRateLimiter(t *testing.T) {
	rl := ratelimiter.NewRateLimiter(2)

	client := "client-1"

	if !rl.Allow(client) {
		t.Fatalf("expected first request to pass")
	}

	if !rl.Allow(client) {
		t.Fatalf("expected second request to pass")
	}

	if rl.Allow(client) {
		t.Fatalf("expected third request to be rate limited")
	}

	time.Sleep(1200 * time.Millisecond)

	if !rl.Allow(client) {
		t.Fatalf("expected token refill")
	}
}
