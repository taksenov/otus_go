package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("check succes", func(t *testing.T) {
		start := time.Now()
		result := RunCmd([]string{"sleep", "1"}, map[string]EnvValue{"LOL": {Value: "KEK", NeedRemove: false}})
		elapsedTime := time.Since(start)

		require.GreaterOrEqual(t, int64(elapsedTime), int64(1000000000))
		require.Equal(t, 0, result)
	})
}
