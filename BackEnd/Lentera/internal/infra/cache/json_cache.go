
package cache

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

// JSONCache menyimpan nilai sebagai JSON di Redis.
type JSONCache struct {
	rdb *redis.Client
}

func NewJSONCache(rdb *redis.Client) *JSONCache {
	return &JSONCache{rdb: rdb}
}

func (c *JSONCache) Get(ctx context.Context, key string, dst any) (bool, error) {
	b, err := c.rdb.Get(ctx, key).Bytes()
	if errors.Is(err, redis.Nil) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, json.Unmarshal(b, dst)
}

func (c *JSONCache) Set(ctx context.Context, key string, val any, ttl time.Duration) error {
	b, err := json.Marshal(val)
	if err != nil {
		return err
	}
	return c.rdb.Set(ctx, key, b, ttl).Err()
}

func (c *JSONCache) Version(ctx context.Context, ns string) (int64, error) {
	v, err := c.rdb.Get(ctx, ns+":ver").Int64()
	if errors.Is(err, redis.Nil) {
		return 0, nil
	}
	return v, err
}

func (c *JSONCache) Bump(ctx context.Context, ns string) error {
	return c.rdb.Incr(ctx, ns+":ver").Err()
}