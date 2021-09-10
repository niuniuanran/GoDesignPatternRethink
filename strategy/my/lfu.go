package main

import (
	"fmt"
	"sort"
)

type ByAccessCount []*item

func (a ByAccessCount) Len() int           { return len(a) }
func (a ByAccessCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAccessCount) Less(i, j int) bool { return a[i].accessCount < a[j].accessCount }

func evictLFU(c *cache) {
	sort.Sort(ByAccessCount(c.items))
	c.removeFirstItem()
	fmt.Println("Evicting by lfu strtegy")
}
