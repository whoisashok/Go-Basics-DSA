package main

type LRUCacheMap struct {
	capacity int
	cache    map[string]int
	order    []string
}

func NewLRUCacheMap(cap int) LRUCacheMap {
	return LRUCacheMap{
		capacity: cap,
		cache:    make(map[string]int),
		order:    []string{},
	}
}

func (l *LRUCacheMap) Get(key string) int {
	if val, ok := l.cache[key]; ok {
		// Move the key to the end (most recently used)
		for i, k := range l.order {
			if k == key {
				// Remove the key from its current position
				l.order = append(l.order[:i], l.order[i+1:]...)
				// Append the key to the end of the order slice
				l.order = append(l.order, key)
				break
			}
		}
		return val
	}
	return -1
}

func (l *LRUCacheMap) Put(key string, value int) {
	if len(l.cache) >= l.capacity {
		// Evict the least recently used item
		oldest := l.order[0]
		// Remove the oldest item from the cache and order
		delete(l.cache, oldest)
		l.order = l.order[1:]
	}
	l.cache[key] = value
	l.order = append(l.order, key)
}

func LRUCacheMain() {
	cache := NewLRUCacheMap(2)
	cache.Put("A", 1)
	cache.Put("B", 2)
	println(cache.Get("A"))
	cache.Put("C", 3) // Evicts key "B"
	println(cache.Get("B"))
	cache.Put("D", 4) // Evicts key "C"
	println(cache.Get("C"))
	println(cache.Get("D"))
}
