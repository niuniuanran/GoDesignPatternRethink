package main

import "fmt"

type testItem struct {
	operation testItemFunc
	key       string
	value     string
}

type testItemFunc func(c *cache, key string, value string) error

func addTestItem(c *cache, key string, value string) error {
	c.add(key, value)
	return nil
}

func getTestItem(c *cache, key string, value string) error {
	v, got := c.get(key)
	if !got {
		return fmt.Errorf("Failed to get item with key %s", key)
	}

	if v != value {
		return fmt.Errorf("Value did not match for key %s, expected %s, got %s", key, value, v)
	}
}
