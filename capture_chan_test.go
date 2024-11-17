package testutils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	testutils "github.com/a-novel-kit/test-utils"
)

func TestCaptureChan(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	t.Run("SingleValue", func(t *testing.T) {
		testutils.SendChan(channel, "Hello world!")
		testutils.RequireChan(t, channel, func(collect *assert.CollectT, value string) {
			require.Equal(collect, "Hello world!", value)
		})
	})

	t.Run("MultipleValues", func(t *testing.T) {
		testutils.SendChan(channel, "Hello world! Again!")
		testutils.RequireChan(t, channel, func(collect *assert.CollectT, value string) {
			require.Equal(collect, "Hello world! Again!", value)
		})
	})
}
