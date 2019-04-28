package cache

import (
	"log"
	"time"

	memoryCache "github.com/patrickmn/go-cache"
)

type Memory struct {
	Expiration time.Duration
	cache      *memoryCache.Cache
}

func (m *Memory) Connect() {
	m.cache = memoryCache.New(time.Duration(m.Expiration)*time.Minute, time.Duration(m.Expiration)*time.Minute)
	log.Println("Using in-memory cache")
}

func (m *Memory) Get(key string) (interface{}, bool) {
	return m.cache.Get(key)
}

func (m *Memory) Set(key string, value interface{}) {
	m.cache.Set(key, value, time.Duration(m.Expiration)*time.Minute)
}
