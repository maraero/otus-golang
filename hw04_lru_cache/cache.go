package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	// Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	Key   Key
	Value interface{}
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if val, ok := c.items[key]; ok {
		return val.Value, ok
	}

	return nil, false
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	if _, exists := c.Get(key); exists {
		c.items[key].Value = value
		c.queue.MoveToFront(c.items[key])
		return true
	}

	c.items[key] = &ListItem{Value: value}
	c.queue.PushFront(c.items[key])
	return false
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
