// Package hw04lrucache -- Otus Go HW03.
package hw04lrucache

// List -- интерфейс.
type List interface {
	Len() int32
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

// ListItem -- узел двусвязного списка.
type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len  int32
	head *ListItem
	tail *ListItem
}

// Len -- возвращяет длину списка.
func (t *list) Len() int32 {
	return t.len
}

// Front -- возвращает элемент начала списка.
func (t *list) Front() *ListItem {
	return t.head
}

// Back -- возвращает элемент конца списка.
func (t *list) Back() *ListItem {
	return t.tail
}

// PushFront -- добавляет элемент в начало.
func (t *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{Value: v}

	if t.len == int32(0) {
		t.head = newItem
		t.tail = newItem
		t.len++
		return newItem
	}

	newItem.Next = t.head
	t.head.Prev = newItem
	t.head = newItem
	t.len++
	return newItem
}

// PushBack -- добавляет элемент в конец.
func (t *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{Value: v}

	if t.len == int32(0) {
		t.PushFront(v)
		return newItem
	}

	newItem.Prev = t.tail
	t.tail.Next = newItem
	t.tail = newItem
	t.len++
	return newItem
}

// Remove -- удаляет элемент из списка.
func (t *list) Remove(i *ListItem) {
	if i.Prev == nil {
		t.popFront()
		return
	}
	if i.Next == nil {
		t.popBack()
		return
	}

	i.Prev.Next = i.Next
	i.Next.Prev = i.Prev

	t.len--
	t.deleteItem(i)
}

// MoveToFront -- перемещает элемент в начало списка.
func (t *list) MoveToFront(i *ListItem) {
	if t.len == int32(0) {
		return
	}
	if t.len == int32(1) {
		return
	}

	oldValue := i.Value

	t.Remove(i)
	t.PushFront(oldValue)
}

// list -- private methods.
func (t *list) deleteItem(i *ListItem) interface{} {
	v := i.Value
	i.Value = nil
	i.Prev = nil
	i.Next = nil
	return v
}

func (t *list) popFront() interface{} {
	oldHead := t.head
	t.head = t.head.Next

	if t.head != nil {
		t.head.Prev = nil
	} else {
		t.tail = nil
	}

	t.len--
	return t.deleteItem(oldHead)
}

func (t *list) popBack() interface{} {
	if t.len == int32(0) {
		return nil
	}

	oldTail := t.tail
	t.tail = oldTail.Prev

	if t.tail != nil {
		t.tail.Next = nil
	}

	t.len--
	return t.deleteItem(oldTail)
}

// NewList -- конструктор.
func NewList() List {
	return new(list)
}
