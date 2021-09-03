package main

import "time"

type cache struct {
	storage      map[string]*item
	evictionAlgo evictionAlgo
	capacity     int
	maxCapacity  int
}

type item struct {
	value        string
	lastUpdated  time.Time
	lastAccessed time.Time
	accessCount  int
}

func initCache(e evictionAlgo) *cache {
	storage := make(map[string]*item)
	return &cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *cache) setEvictionAlgo(e evictionAlgo) {
	c.evictionAlgo = e
}

func (c *cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = &item{
		value:       value,
		lastUpdated: time.Now(),
		accessCount: 0}
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
	c.evictionAlgo.evict(c)
	c.capacity--
}
