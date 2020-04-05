/*

References:
- https://github.com/golang/groupcache/blob/master/lru/lru.go

*/

package main

import (
	"container/list"
	"fmt"
)

// Cache implements an LRU cache.
type Cache struct {
	Size  int
	list  *list.List
	cache map[interface{}]*list.Element
}

// entry holds a key and a value for each element in the list.
type entry struct {
	key   interface{}
	value interface{}
}

// New creates a new cache.
func New(size int) *Cache {
	return &Cache{
		Size:  size,
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
	if c.Size != 0 && c.list.Len() > c.Size {
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

func main() {
	c := New(3)
	c.Add(1, 1)
	c.Add(2, 2)
	c.Add(3, 3)
	c.Add(4, 4)
	fmt.Println(c.Get(1))
	fmt.Println(c.Get(2))
	fmt.Println(c.Get(3))
	fmt.Println(c.Get(4))
	fmt.Println(c.Get(5))
}
