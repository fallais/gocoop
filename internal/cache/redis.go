package cache

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// RedisCache is a Redis cache
type RedisCache struct {
	client            *redis.Client
	defaultExpiration time.Duration
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewRedisCache returns a new RedisCache
func NewRedisCache(host, password string, defaultExpiration time.Duration) (*RedisCache, error) {
	// Create the client
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
	})

	// Check if the Redis responds
	_, err := client.Ping().Result()
	if err != nil {
		return nil, fmt.Errorf("Redis cache is unavailable")
	}

	return &RedisCache{
		client:            client,
		defaultExpiration: defaultExpiration,
	}, nil
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// Set a value in Redis cache
func (c *RedisCache) Set(key string, value interface{}, expires time.Duration) error {
	return c.client.Set(key, value, expires).Err()
}

// Get a value from Redis cache
func (c *RedisCache) Get(key string) (string, error) {
	return c.client.Get(key).Result()
}

// Exists a value from Redis cache
func (c *RedisCache) Exists(key string) (string, error) {
	return c.client.Get(key).Result()
}

// Delete a value in Redis cache
func (c *RedisCache) Delete(key string) error {
	_, err := c.client.Del(key).Result()

	return err
}

// Flush the Redis cache
func (c *RedisCache) Flush() error {
	return c.client.FlushAll().Err()
}
