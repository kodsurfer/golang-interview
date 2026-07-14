package ratelimiter

import (
	"sync"
	"time"
)

type RateLimiter struct {
	capacity   float64
	tokens     float64
	lastUpdate time.Time
	mu         sync.Mutex
}

func NewRateLimiter(rps float64, burst float64) *RateLimiter {
	return &RateLimiter{
		capacity:   rps,
		tokens:     burst,
		lastUpdate: time.Now(),
	}

}

func (r *RateLimiter) Allow() bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(r.lastUpdate).Seconds()
	r.lastUpdate = now

	r.tokens += elapsed * r.tokens
	if r.tokens > r.capacity {
		r.tokens = r.capacity
	}

	if r.tokens >= 1 {
		r.tokens -= 1
		return true
	}

	return false
}
