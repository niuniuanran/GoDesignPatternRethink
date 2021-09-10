package main

import (
	"sort"
)

type ByAccessTime []*item

func (a ByAccessTime) Len() int           { return len(a) }
func (a ByAccessTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAccessTime) Less(i, j int) bool { return a[i].lastAccessed.Before(a[j].lastAccessed) }

func evictLRU(c *cache) {
	sort.Sort(ByAccessTime(c.items))
	c.removeFirstItem()
}
