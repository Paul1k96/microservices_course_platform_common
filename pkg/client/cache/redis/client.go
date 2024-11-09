package redis

import (
	"context"
	"log"
	"time"

	"github.com/Paul1k96/microservices_course_platform_common/pkg/client/cache"
	"github.com/gomodule/redigo/redis"
)

var _ cache.RedisClient = (*Client)(nil)

type handler func(ctx context.Context, conn redis.Conn) error

// Client represents redis client.
type Client struct {
	pool   *redis.Pool
	config Config
}

// NewClient creates a new instance of redis client.
func NewClient(pool *redis.Pool, config Config) *Client {
	return &Client{
		pool:   pool,
		config: config,
	}
}

// HSet sets hash values.
func (c *Client) HSet(ctx context.Context, key string, values interface{}) error {
	err := c.execute(ctx, func(_ context.Context, conn redis.Conn) error {
		_, err := conn.Do("HSET", redis.Args{key}.AddFlat(values)...)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// Set sets value.
func (c *Client) Set(ctx context.Context, key string, value interface{}) error {
	err := c.execute(ctx, func(_ context.Context, conn redis.Conn) error {
		_, err := conn.Do("SET", redis.Args{key}.Add(value)...)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// HGetAll gets all hash values.
func (c *Client) HGetAll(ctx context.Context, key string) ([]interface{}, error) {
	var values []interface{}
	err := c.execute(ctx, func(_ context.Context, conn redis.Conn) error {
		var errEx error
		values, errEx = redis.Values(conn.Do("HGETALL", key))
		if errEx != nil {
			return errEx
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}

// Get gets value.
func (c *Client) Get(ctx context.Context, key string) (interface{}, error) {
	var value interface{}
	err := c.execute(ctx, func(_ context.Context, conn redis.Conn) error {
		var errEx error
		value, errEx = conn.Do("GET", key)
		if errEx != nil {
			return errEx
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return value, nil
}

// Delete deletes key.
func (c *Client) Delete(ctx context.Context, key string) error {
	err := c.execute(ctx, func(_ context.Context, conn redis.Conn) error {
		_, err := conn.Do("DEL", key)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// Expire sets expiration time for key.
func (c *Client) Expire(ctx context.Context, key string, expiration time.Duration) error {
	err := c.execute(ctx, func(_ context.Context, conn redis.Conn) error {
		_, err := conn.Do("EXPIRE", key, int(expiration.Seconds()))
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// Ping checks connection.
func (c *Client) Ping(ctx context.Context) error {
	err := c.execute(ctx, func(_ context.Context, conn redis.Conn) error {
		_, err := conn.Do("PING")
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) execute(ctx context.Context, handler handler) error {
	conn, err := c.getConnect(ctx)
	if err != nil {
		return err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Printf("failed to close redis connection: %v\n", err)
		}
	}()

	err = handler(ctx, conn)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) getConnect(ctx context.Context) (redis.Conn, error) {
	getConnTimeoutCtx, cancel := context.WithTimeout(ctx, c.config.GetConnectionTimeout())
	defer cancel()

	conn, err := c.pool.GetContext(getConnTimeoutCtx)
	if err != nil {
		log.Printf("failed to get redis connection: %v\n", err)

		_ = conn.Close()
		return nil, err
	}

	return conn, nil
}
