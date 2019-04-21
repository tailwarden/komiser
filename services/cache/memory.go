package cache

import (
	"time"

	memoryCache "github.com/patrickmn/go-cache"
)

type Memory struct {
	Expiration time.Duration
	cache      *memoryCache.Cache
}

func (m *Memory) Connect() {
	m.cache = memoryCache.New(m.Expiration*time.Minute, m.Expiration*time.Minute)
}

func (m *Memory) Get(key string) (interface{}, bool) {
	return m.cache.Get(key)
}

func (m *Memory) Set(key string, value interface{}) {
	m.cache.Set(key, value, m.Expiration)
}
