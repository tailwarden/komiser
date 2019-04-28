package cache

import (
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis"
)

type Redis struct {
	Addr       string
	Expiration time.Duration
	client     *redis.Client
}

func (r *Redis) Connect() {
	r.client = redis.NewClient(&redis.Options{
		Addr: r.Addr,
		DB:   0,
	})
	_, err := r.client.Ping().Result()
	if err != nil {
		log.Fatal("Cloudn't connect to Redis:", err)
	} else {
		log.Println("Successfully connected to Redis")
	}
}

func (r *Redis) Get(key string) (interface{}, bool) {
	val, err := r.client.Get(key).Result()
	if err == redis.Nil {
		return val, false
	}
	var data interface{}
	json.Unmarshal([]byte(val), &data)
	return data, true
}

func (r *Redis) Set(key string, value interface{}) {
	data, err := json.Marshal(value)
	if err != nil {
		log.Fatal(err)
	}
	err = r.client.Set(key, data, r.Expiration*time.Minute).Err()
	if err != nil {
		log.Fatal(err)
	}
}
