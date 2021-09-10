package main

import (
	"sort"
)

type ByUpdateTime []*item

func (a ByUpdateTime) Len() int           { return len(a) }
func (a ByUpdateTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByUpdateTime) Less(i, j int) bool { return a[i].lastUpdated.Before(a[j].lastUpdated) }

func evictFIFO(c *cache) {
	sort.Sort(ByUpdateTime(c.items))
	c.removeFirstItem()
}
