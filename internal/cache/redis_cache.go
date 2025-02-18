package cache

import (
	"chat-server/config"
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

func InitRedis(env *config.Env) *RedisCache {
	addr := env.RedisHost + ":" + env.RedisPort
	password := env.RedisPass
	db := env.RedisName

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	// Check connection
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Println("❌ Redis connection failed:", err)
		panic(err)
	}

	log.Println("✅ Connected to Redis")

	return &RedisCache{client: rdb}
}

func (r *RedisCache) Get(key string) (string, error) {
	ctx := context.Background()
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return val, nil
}

func (r *RedisCache) Set(key string, value string, expiration time.Duration) error {
	ctx := context.Background()
	return r.client.Set(ctx, key, value, expiration).Err()
}

func (r *RedisCache) Delete(key string) error {
	ctx := context.Background()
	return r.client.Del(ctx, key).Err()
}

func (r *RedisCache) Close() error {
	return r.client.Close()
}
