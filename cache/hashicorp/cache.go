package hashicorp

import (
	"fmt"

	"github.com/hashicorp/golang-lru"
)

const (
	ARC      = "arc"
	LRU      = "lru"
	TwoQueue = "2q"
)

// LRUCache is the interface for simple LRU cache.
type CacheInterface interface {
	// Adds a value to the cache, returns true if an eviction occurred and
	// updates the "recently used"-ness of the key.
	Add(key, value interface{})

	// Returns key's value from the cache and
	// updates the "recently used"-ness of the key. #value, isFound
	Get(key interface{}) (value interface{}, ok bool)

	// Check if a key exsists in cache without updating the recent-ness.
	Contains(key interface{}) (ok bool)

	// Returns key's value without updating the "recently used"-ness of the key.
	Peek(key interface{}) (value interface{}, ok bool)

	// Removes a key from the cache.
	Remove(key interface{})

	// Returns a slice of the keys in the cache, from oldest to newest.
	Keys() []interface{}

	// Returns the number of items in the cache.
	Len() int

	// Clear all cache entries
	Purge()
}

func NewCache(algorithm string, size int) (CacheInterface, error) {
	switch algorithm {
	case LRU:
		return NewLRU(size)
	case ARC:
		return lru.NewARC(size)
	case TwoQueue:
		return lru.New2Q(size)
	default:
		return nil, fmt.Errorf("not support cache algorithm %s", algorithm)
	}
}

type LRUCache struct {
	lru.Cache
}

func NewLRU(size int) (*LRUCache, error) {
	cache, err := lru.New(size)
	if err != nil {
		return nil, err
	}
	return &LRUCache{*cache}, nil
}

func (lru *LRUCache) Add(key, value interface{}) {
	lru.Cache.Add(key, value)
}
