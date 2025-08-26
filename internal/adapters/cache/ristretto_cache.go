package cache

import (
	"github.com/dgraph-io/ristretto"
	"log"
)

type RistrettoCache struct {
	c *ristretto.Cache
}

func NewRistrettoCache() *RistrettoCache {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,
		MaxCost:     1 << 30,
		BufferItems: 64,
	})
	if err != nil {
		log.Fatalf("failed to create ristretto cache: %v", err)
	}

	return &RistrettoCache{c: cache}
}

func (r *RistrettoCache) Set(key string, value interface{}, cost int64) {
	r.c.Set(key, value, cost)
	r.c.Wait()
}

func (r *RistrettoCache) Get(key string) (interface{}, bool) {
	return r.c.Get(key)
}

func (r *RistrettoCache) Delete(key string) {
	r.c.Del(key)
}
