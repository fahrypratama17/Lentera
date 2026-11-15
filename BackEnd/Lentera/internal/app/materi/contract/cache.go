package contract

import (
	"context"
	"time"
)

type Cache interface {
	// Get mengisi dst dari cache. ok=false kalau tidak ada (miss).
	Get(ctx context.Context, key string, dst any) (ok bool, err error)
	Set(ctx context.Context, key string, val any, ttl time.Duration) error
	// Version & Bump untuk invalidasi: semua key yang memakai versi lama otomatis tidak terpakai.
	Version(ctx context.Context, ns string) (int64, error)
	Bump(ctx context.Context, ns string) error
}

// NoopCache: dipakai kalau Redis tidak tersedia & di unit test. Selalu miss.
type NoopCache struct{}

func (NoopCache) Get(context.Context, string, any) (bool, error)             { return false, nil }
func (NoopCache) Set(context.Context, string, any, time.Duration) error      { return nil }
func (NoopCache) Version(context.Context, string) (int64, error)             { return 0, nil }
func (NoopCache) Bump(context.Context, string) error                         { return nil }