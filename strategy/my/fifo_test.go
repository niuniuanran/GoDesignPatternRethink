package main

import "testing"

func TestFifoEvict(t *testing.T) {
	var fifoTests = []struct {
		index       int
		name        string
		maxCapacity int
		testItems   []testItem
		gotKeys     []string
		notGotKeys  []string
	}{
		{1, "First in first out", 3, []testItem{
			{addTestItem, "a", "1"},
			{addTestItem, "b", "2"},
			{addTestItem, "c", "3"},
			{addTestItem, "d", "4"},
		}, []string{"b", "c", "d"}, []string{"a"}},
		{2, "Access frequency should have effect", 3, []testItem{
			{addTestItem, "a", "1"},
			{addTestItem, "b", "2"},
			{getTestItem, "a", "1"},
			{getTestItem, "a", "1"},
			{addTestItem, "c", "3"},
			{addTestItem, "d", "4"},
		}, []string{"b", "c", "d"}, []string{"a"}},
		{3, "Access time should have no effect", 3, []testItem{
			{addTestItem, "a", "1"},
			{addTestItem, "b", "2"},
			{getTestItem, "b", "2"},
			{getTestItem, "a", "1"},
			{getTestItem, "a", "1"},
			{addTestItem, "c", "3"},
			{addTestItem, "d", "4"},
		}, []string{"b", "c", "d"}, []string{"a"}},
	}

	for _, tt := range fifoTests {
		cache := initCache(evictFIFO, tt.maxCapacity)
		for _, item := range tt.testItems {
			err := item.operation(cache, item.key, item.value)
			if err != nil {
				t.Errorf("Test %d - %s: %s", tt.index, tt.name, err.Error())
			}
		}
		for _, k := range tt.gotKeys {
			_, got := cache.get(k)
			if !got {
				t.Errorf("Test %d - %s: Expected item in cache is evected, key: %s", tt.index, tt.name, k)
			}
		}

		for _, k := range tt.notGotKeys {
			_, got := cache.get(k)
			if got {
				t.Errorf("Test %d - %s: Item expected to be evicted but still in cache, key: %s", tt.index, tt.name, k)
			}
		}
	}

}
