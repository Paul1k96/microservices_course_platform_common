package cache

import (
	"context"
	"time"
)

//go:generate ../../../bin/mockgen -source $GOFILE -destination "mocks/cache.go" -package mocks

// RedisClient represents Redis client.
type RedisClient interface {
	HSet(ctx context.Context, key string, values interface{}) error
	Set(ctx context.Context, key string, value interface{}) error
	HGetAll(ctx context.Context, key string) ([]interface{}, error)
	Get(ctx context.Context, key string) (interface{}, error)
	Delete(ctx context.Context, key string) error
	Expire(ctx context.Context, key string, expiration time.Duration) error
	Ping(ctx context.Context) error
}
