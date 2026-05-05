package main

import (
	"sync"
	"time"
)

type tokenbucket struct {
	tokens    float64
	maxtokens float64
	rate      float64
	lasttime  time.Time
	mu        sync.Mutex
}

var buckets sync.Map

func getbucket(ip string) *tokenbucket {
	if val, ok := buckets.Load(ip); ok {
		return val.(*tokenbucket)
	}
	b := &tokenbucket{
		tokens:    float64(config.rateburst),
		maxtokens: float64(config.rateburst),
		rate:      float64(config.ratepersec),
		lasttime:  time.Now(),
	}
	buckets.Store(ip, b)
	return b
}

func allowrequest(ip string) bool {
	b := getbucket(ip)
	b.mu.Lock()
	defer b.mu.Unlock()
	now := time.Now()
	elapsed := now.Sub(b.lasttime).Seconds()
	b.tokens += elapsed * b.rate
	if b.tokens > b.maxtokens {
		b.tokens = b.maxtokens
	}
	b.lasttime = now
	if b.tokens >= 1.0 {
		b.tokens--
		return true
	}
	return false
}
