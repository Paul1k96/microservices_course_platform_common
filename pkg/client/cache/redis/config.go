package redis

import "time"

// Config represents configuration for Redis.
type Config interface {
	GetAddress() string
	GetConnectionTimeout() time.Duration
	GetMaxIdle() int
	GetIdleTimeout() time.Duration
}
