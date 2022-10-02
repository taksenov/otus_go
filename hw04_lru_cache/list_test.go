package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func makeValuesSlice(l List) []int32 {
	elems := make([]int32, 0, l.Len())
	for i := l.Front(); i != nil; i = i.Next {
		val, _ := i.Value.(int32)
		elems = append(elems, val)
	}
	return elems
}

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, int32(0), l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(int32(10)) // [10]
		l.Front()
		require.Equal(t, []int32{10}, makeValuesSlice(l))

		l.PushBack(int32(20)) // [10, 20]
		require.Equal(t, []int32{10, 20}, makeValuesSlice(l))

		l.PushBack(int32(30)) // [10, 20, 30]
		require.Equal(t, []int32{10, 20, 30}, makeValuesSlice(l))
		require.Equal(t, int32(3), l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, int32(2), l.Len())
		require.Equal(t, []int32{10, 30}, makeValuesSlice(l))

		for i, v := range [...]int32{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]
		require.Equal(t, int32(7), l.Len())
		require.Equal(t, int32(80), l.Front().Value)
		require.Equal(t, int32(70), l.Back().Value)
		require.Equal(t, []int32{80, 60, 40, 10, 30, 50, 70}, makeValuesSlice(l))

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		require.Equal(t, []int32{80, 60, 40, 10, 30, 50, 70}, makeValuesSlice(l))

		l.MoveToFront(l.Back()) // [70, 80, 60, 40, 10, 30, 50]
		require.Equal(t, []int32{70, 80, 60, 40, 10, 30, 50}, makeValuesSlice(l))
	})
}
