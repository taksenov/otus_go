package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMasterOfChanks(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		result := masterOfChunks(1000)
		require.Equal(t, int64(10), result.count)
		require.Equal(t, int64(100), result.size)
		require.Equal(t, int64(0), result.tail)
	})

	t.Run("simple two", func(t *testing.T) {
		result := masterOfChunks(1223)
		require.Equal(t, int64(10), result.count)
		require.Equal(t, int64(122), result.size)
		require.Equal(t, int64(3), result.tail)
	})

	t.Run("less than measure", func(t *testing.T) {
		result := masterOfChunks(9)
		require.Equal(t, int64(1), result.count)
		require.Equal(t, int64(9), result.size)
		require.Equal(t, int64(0), result.tail)
	})

	t.Run("zero", func(t *testing.T) {
		result := masterOfChunks(0)
		require.Equal(t, int64(0), result.count)
		require.Equal(t, int64(0), result.size)
		require.Equal(t, int64(0), result.tail)
	})

	t.Run("negative", func(t *testing.T) {
		result := masterOfChunks(-1)
		require.Equal(t, int64(0), result.count)
		require.Equal(t, int64(0), result.size)
		require.Equal(t, int64(0), result.tail)
	})
}
