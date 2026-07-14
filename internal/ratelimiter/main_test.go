package ratelimiter

import (
	"sync"
	"testing"
)

func TestRateLimiter_Concurrent(t *testing.T) {
	limiter := NewRateLimiter(10, 20)
	const total = 1000
	allowed := 0
	var wg sync.WaitGroup

	for i := 0; i < total; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if limiter.Allow() {
				allowed++
			}
		}()
	}
	wg.Wait()

	if allowed == total {
		t.Errorf("expected some tokens rejected, got all allowed")
	}
}
