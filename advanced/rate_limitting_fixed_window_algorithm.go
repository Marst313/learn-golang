package main

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiterFixedWindow struct {
	mu        sync.Mutex
	count     int
	limit     int
	window    time.Duration // LIMIT BERAAPA LAMA ??
	resetTime time.Time
}

func NewRateLimiterFixedWindow(limit int, window time.Duration) *RateLimiterFixedWindow {
	return &RateLimiterFixedWindow{
		limit:  limit,
		window: window,
	}
}
func (rl *RateLimiterFixedWindow) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	now := time.Now()

	if now.After(rl.resetTime) {
		rl.resetTime = now.Add(rl.window)
		rl.count = 0
	}

	// CHECK IF THE LIMIT IS NOT OVER YET
	if rl.count < rl.limit {
		rl.count++
		return true
	}

	return false
}

func rate_limitting_fixed_window_algorithm() {
	// var wg sync.WaitGroup
	rateLimiter := NewRateLimiterFixedWindow(10, 5*time.Second)

	for i := range 100 {
		// wg.Add(1)
		// go func() {
		if rateLimiter.Allow() {
			fmt.Println("Request Allowed.", i+1)
		} else {
			fmt.Println("Request denied", time.Now())
		}
		// wg.Done()
		// }()
		time.Sleep(200 * time.Millisecond)
	}
	// wg.Wait()

}
