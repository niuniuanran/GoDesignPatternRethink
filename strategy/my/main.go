package main

func main() {
	cache := initCache(evictFIFO)
	cache.add("a", "1")
	cache.add("b", "2")
	cache.get("a")
	cache.get("a")
	cache.get("a")
	cache.get("b")
	cache.add("c", "3")
	cache.add("d", "3")
	cache.printItemsVerbose()

	cache.clear()
	cache.setEvictionFunc(evictLFU)
	cache.add("a", "1")
	cache.add("b", "2")
	cache.get("a")
	cache.get("a")
	cache.get("a")
	cache.get("b")
	cache.add("c", "3")
	cache.printItemsVerbose()

	cache.clear()
	cache.setEvictionFunc(evictLRU)
	cache.add("a", "1")
	cache.add("b", "2")
	cache.get("a")
	cache.get("a")
	cache.get("a")
	cache.get("b")
	cache.add("c", "3")
	cache.printItemsVerbose()

}
