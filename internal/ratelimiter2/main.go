package ratelimiter2

import "time"

type RateLimiter2 struct {
	ticker *time.Ticker
	tokens chan struct{}
	stopCh chan struct{}
}

func (r *RateLimiter2) refill() {
	for {
		select {
		case <-r.ticker.C:
			select {
			case r.tokens <- struct{}{}:
			default:
			}
		case <-r.stopCh:
			r.ticker.Stop()
			return
		}
	}
}

func (r *RateLimiter2) Allow() bool {
	select {
	case <-r.tokens:
		return true
	default:
		return false
	}
}

func (r RateLimiter2) Stop() {
	close(r.stopCh)
}

func NewRateLimiter2(rate int, dur time.Duration) *RateLimiter2 {
	r := &RateLimiter2{
		ticker: time.NewTicker(dur),
		tokens: make(chan struct{}, rate),
		stopCh: make(chan struct{}),
	}

	for i := 0; i < rate; i++ {
		r.tokens <- struct{}{}
	}

	go r.refill()

	return r
}
