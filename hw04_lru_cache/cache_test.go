package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(1)

		c.Set("aaa", "bbb")
		c.Clear()
		val, ok := c.Get("aaa")

		require.False(t, ok)
		require.Equal(t, nil, val)
	})

	t.Run("pushing elements", func(t *testing.T) {
		c := NewCache(3)

		c.Set("a", "LOL")
		v, ok := c.Get("a")
		require.True(t, ok)
		require.Equal(t, "LOL", v)

		c.Set("b", "KEK")
		v, ok = c.Get("b")
		require.True(t, ok)
		require.Equal(t, "KEK", v)

		c.Set("c", "AZAZA")
		v, ok = c.Get("c")
		require.True(t, ok)
		require.Equal(t, "AZAZA", v)

		c.Set("d", "is palindromes")
		v, ok = c.Get("a")
		require.False(t, ok)
		require.Equal(t, nil, v)
	})

	t.Run("pushing old elements", func(t *testing.T) {
		c := NewCache(3)

		c.Set("a", "LOL")
		c.Set("b", "KEK")
		c.Set("old", "AZAZA")

		c.Get("old")
		c.Set("old", "OLOLO")
		c.Get("b")
		c.Set("b", "LOL")
		c.Get("a")
		c.Set("a", "KEK")
		c.Get("b")
		c.Get("a")
		c.Set("d", "is palindromes")

		v, ok := c.Get("old")
		require.False(t, ok)
		require.Equal(t, nil, v)
	})
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
