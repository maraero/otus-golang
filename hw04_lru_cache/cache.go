package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
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
	if c.items[key] == nil {
		return nil, false
	}

	if actualVal, ok := c.items[key].Value.(cacheItem); ok {
		return actualVal.Value, ok
	}

	return nil, false
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	if _, exists := c.items[key]; exists {
		item := c.items[key]
		item.Value = cacheItem{Value: value, Key: key}
		c.queue.MoveToFront(item)
		return true
	}

	if c.queue.Len() == c.capacity {
		lastItem := c.queue.Back()
		c.queue.Remove(lastItem)

		valWithKey, ok := lastItem.Value.(cacheItem)

		if ok {
			delete(c.items, valWithKey.Key)
		}
	}

	item := c.queue.PushFront(cacheItem{Value: value, Key: key})
	c.items[key] = item
	return false
}

func (c *lruCache) Clear() {
	c.items = make(map[Key]*ListItem, c.capacity)
	c.queue = NewList()
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
