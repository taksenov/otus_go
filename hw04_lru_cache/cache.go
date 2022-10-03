// Package hw04lrucache -- Otus Go HW03.
package hw04lrucache

// Key -- ключ-строка.
type Key string

// Cache -- интерфейс.
type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	Cache // Remove me after realization.

	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

// Get -- получает значение из кеша по ключу.
func (t *lruCache) Get(k Key) (interface{}, bool) {
	// @todo реализовать функциональность
	return nil, false
}

// Set -- получает значение из кеша по ключу.
func (t *lruCache) Set(k Key, v interface{}) bool {
	// @todo реализовать функциональность
	return false
}

// Clear -- очищает кеш.
func (t *lruCache) Clear() {
	// @todo реализовать функциональность
	return
}

// NewCache -- конструктор.
func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
