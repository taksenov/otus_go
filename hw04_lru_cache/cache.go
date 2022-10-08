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
	capacity int32
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

// Get -- получает значение из кеша по ключу.
func (t *lruCache) Get(k Key) (interface{}, bool) {
	elem, ok := t.items[k]
	if ok {
		v, _ := elem.Value.(*cacheItem)
		t.queue.MoveToFront(elem)

		return v.value, ok
	}

	return nil, ok
}

// Set -- устанавливает значение в кеш по ключу.
func (t *lruCache) Set(k Key, v interface{}) bool {
	cItem := &cacheItem{key: k, value: v}
	elem, ok := t.items[k]
	if ok {
		elem.Value = cItem
		t.items[k] = elem
		t.queue.MoveToFront(elem)

		return ok
	}

	t.items[k] = t.queue.PushFront(cItem)

	if t.queue.Len() > t.capacity {
		lastEl := t.queue.Back()
		t.queue.Remove(lastEl)
		lK := lastEl.Value.(*cacheItem)
		delete(t.items, lK.key)
	}

	return ok
}

// Clear -- очищает кеш.
func (t *lruCache) Clear() {
	if t == nil {
		return
	}
	if t.capacity == int32(0) {
		return
	}

	c := t.capacity
	t.queue = NewList()
	t.items = make(map[Key]*ListItem, c)
}

// NewCache -- конструктор.
func NewCache(capacity int32) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
