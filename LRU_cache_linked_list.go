package main

import (
	"container/list"
	"fmt"
)

// Cache is a Least Recently Used (LRU) cache.
type LRUCache struct {
	capacity int
	// map stores key to a list element pointer for O(1) lookup
	items map[any]*list.Element
	// list is a doubly linked list that maintains the order of usage (MRU at front, LRU at back)
	evictionList *list.List
}

// entry is a struct stored in the list element.
type entry struct {
	key   any
	value any
}

// NewLRUCache creates a new LRUCache with the given capacity.
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity:     capacity,
		items:        make(map[any]*list.Element),
		evictionList: list.New(),
	}
}

// Get retrieves a value from the cache and marks it as recently used.
func (c *LRUCache) Get(key interface{}) (value interface{}, ok bool) {
	if element, found := c.items[key]; found {
		// Move the element to the front of the list (most recently used)
		c.evictionList.MoveToFront(element)
		return element.Value.(*entry).value, true
	}
	return nil, false
}

// Put adds a value to the cache, or updates the value if the key already exists.
// It handles capacity limits by evicting the least recently used item.
func (c *LRUCache) Put(key interface{}, value interface{}) {
	if element, found := c.items[key]; found {
		// Key exists, update value and move to front
		c.evictionList.MoveToFront(element)
		element.Value.(*entry).value = value
		return
	}

	// Key is new, add as a new element to the front
	newEntry := &entry{key: key, value: value}
	element := c.evictionList.PushFront(newEntry)
	c.items[key] = element

	// Check capacity and evict LRU if full
	if c.evictionList.Len() > c.capacity {
		c.evictLeastRecentlyUsed()
	}
}

// evictLeastRecentlyUsed removes the element at the tail of the list (LRU element).
func (c *LRUCache) evictLeastRecentlyUsed() {
	if c.evictionList.Len() == 0 {
		return
	}
	// Remove the tail element from the list
	lruElement := c.evictionList.Back()
	if lruElement != nil {
		c.evictionList.Remove(lruElement)
		// Also remove from the map
		lruKey := lruElement.Value.(*entry).key
		delete(c.items, lruKey)
	}
}

func main1() {
	cache := NewLRUCache(3)

	cache.Put("A", 1)
	cache.Put("B", 2)
	cache.Put("C", 3)

	fmt.Printf("Cache size: %d\n", cache.evictionList.Len()) // Output: 3

	// Access "A" to make it Most Recently Used (MRU)
	if val, ok := cache.Get("A"); ok {
		fmt.Printf("Get A: %v\n", val) // Output: Get A: 1
	}

	// Put "D". This exceeds capacity, so "B" (LRU) should be evicted.
	cache.Put("D", 4)

	fmt.Printf("Cache size after adding D: %d\n", cache.evictionList.Len()) // Output: 3

	// Check if B is still in cache
	if _, ok := cache.Get("B"); !ok {
		fmt.Println("Get B: Not Found (correctly evicted)")
	}

	// Check if C is in cache (it should be)
	if val, ok := cache.Get("C"); ok {
		fmt.Printf("Get C: %v\n", val) // Output: Get C: 3
	}
}
