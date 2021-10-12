package hw04lrucache

type Key string

type Cache interface {
	// Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	// Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

// type cacheItem struct {
// 	key   Key
// 	value interface{}
// }

func (c *lruCache) Get(key Key) (interface{}, bool) {
	val, ok := c.items[key]

	return val, ok
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
