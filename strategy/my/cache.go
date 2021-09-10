package main

import (
	"fmt"
	"time"
)

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

func initCache(e func(c *cache), maxCapacity int) *cache {
	storage := make(map[string]*item)
	return &cache{
		storage:     storage,
		items:       make([]*item, 0),
		evictFunc:   e,
		capacity:    0,
		maxCapacity: maxCapacity,
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
}

func (c *cache) printItemsVerbose() {
	fmt.Println("Key\tValue\tAccess Count\tLast Accessed\t\t\t\tLast Updated\t\t")
	for _, i := range c.items {
		fmt.Printf("%s\t%s\t%d\t\t%s\t%s\t\n", i.key, i.value, i.accessCount, i.lastAccessed, i.lastUpdated)
	}
}

func (c *cache) printItems() {
	for _, i := range c.items {
		fmt.Print(i.key, " ")
	}
	fmt.Println()
}

func (c *cache) removeFirstItem() {
	toEvict := c.items[0]
	delete(c.storage, toEvict.key)
	c.items = c.items[1:]
	c.capacity--
}

func (c *cache) clear() {
	for k := range c.storage {
		delete(c.storage, k)
	}
	c.items = make([]*item, 0)
	c.capacity = 0
}
