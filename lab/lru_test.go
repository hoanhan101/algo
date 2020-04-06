/*
Problem:
- Implement LRU cache implementation using a doubly linked list and a hashmap.
  The cache should have these following methods:
  - Set: put a key-value pair in the cache
  - Get: get a key's value
  - Len: return a number of items in the cache
  - Remove: remove a key-value pair
  - Clear: purge the cache

Cost:
- O(1) time, O(n) space.

References:
- https://github.com/golang/groupcache/blob/master/lru/lru.go
- https://www.interviewcake.com/concept/java/lru-cache
*/

package lab

import (
	"container/list"
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestCache(t *testing.T) {
	c := NewCache(3)

	common.Equal(t, 0, c.Len())

	for i := 0; i < 6; i++ {
		c.Add(i, i)
	}

	common.Equal(t, 3, c.Len())

	for i := 0; i < 3; i++ {
		common.Equal(t, nil, c.Get(i))
	}

	for i := 3; i < 6; i++ {
		common.Equal(t, i, c.Get(i))
	}

	c.Clear()

	common.Equal(t, 0, c.Len())

	for i := 3; i < 6; i++ {
		common.Equal(t, nil, c.Get(i))
	}

	c.Add(6, 6)
	common.Equal(t, 6, c.Get(6))
	c.Remove(6)
	common.Equal(t, nil, c.Get(6))
	common.Equal(t, 0, c.Len())
}

// Cache implements an LRU cache.
type Cache struct {
	size  int
	list  *list.List
	cache map[interface{}]*list.Element
}

// entry holds a key and a value for each element in the list.
type entry struct {
	key   interface{}
	value interface{}
}

// New creates a new cache.
func NewCache(size int) *Cache {
	return &Cache{
		size:  size,
		list:  list.New(),
		cache: make(map[interface{}]*list.Element),
	}
}

// Add adds a value to the cache.
func (c *Cache) Add(key, value interface{}) {
	// make sure if the cache is not nil.
	if c.cache == nil {
		c.list = list.New()
		c.cache = make(map[interface{}]*list.Element)
	}

	// if the item is found, move it to front of the list and update its value.
	if e, ok := c.cache[key]; ok {
		c.list.MoveToFront(e)
		e.Value.(*entry).value = value
		return
	}

	// else, create a new entry, push it to the front of the list and put it in
	// the cache.
	en := c.list.PushFront(&entry{key, value})
	c.cache[key] = en

	// if the cache is full, evict the least-frequently used item at the back of
	// the list.
	if c.size != 0 && c.list.Len() > c.size {
		c.remove(c.list.Back())
	}
}

// Get looks up a key's value from the cache.
func (c *Cache) Get(key interface{}) interface{} {
	// return nil immediately if the cache is nil.
	if c.cache == nil {
		return nil
	}

	// if the item is found, move it to front of the list and return its value.
	if e, ok := c.cache[key]; ok {
		c.list.MoveToFront(e)
		return e.Value.(*entry).value
	}

	return nil
}

// Remove removes the provided key from the cache.
func (c *Cache) Remove(key interface{}) {
	if c.cache == nil {
		return
	}

	if e, ok := c.cache[key]; ok {
		c.remove(e)
	}
}

// Len returns the number of items in the cache.
func (c *Cache) Len() int {
	if c.cache == nil {
		return 0
	}

	return c.list.Len()
}

// Clear purges all stored items from the cache.
func (c *Cache) Clear() {
	c.list = nil
	c.cache = nil
}

// remove deletes the element in the list as well as the cache.
func (c *Cache) remove(e *list.Element) {
	c.list.Remove(e)
	delete(c.cache, e.Value.(*entry).key)
}
