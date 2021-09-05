package main

import "time"

type evictionFunc func(c *cache)

type cache struct {
	storage     map[string]*item
	items       []*item
	evictFunc   evictionFunc
	capacity    int
	maxCapacity int
}

type item struct {
	key          string
	value        string
	lastUpdated  time.Time
	lastAccessed time.Time
	accessCount  int
}

func initCache(e func(c *cache)) *cache {
	storage := make(map[string]*item)
	return &cache{
		storage:     storage,
		items:       make([]*item, 0),
		evictFunc:   e,
		capacity:    0,
		maxCapacity: 2,
	}
}

func (c *cache) setEvictionFunc(e evictionFunc) {
	c.evictFunc = e
}

func (c *cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = &item{
		key:         key,
		value:       value,
		lastUpdated: time.Now(),
		accessCount: 0}
	c.items = append(c.items, c.storage[key])
}

func (c *cache) get(key string) (string, bool) {
	item, got := c.storage[key]
	if !got {
		return "", false
	}
	item.lastAccessed = time.Now()
	item.accessCount++
	return item.value, true
}

func (c *cache) evict() {
	c.evictFunc(c)
	c.capacity--
}
