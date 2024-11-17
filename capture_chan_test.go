package testutils_test

import (
	"testing"

	testutils "github.com/a-novel-kit/test-utils"
)

func TestCaptureChan(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	t.Run("SingleValue", func(t *testing.T) {
		testutils.SendChan(channel, "Hello world!")
		testutils.RequireChanEqual(t, channel, "Hello world!")
	})

	t.Run("MultipleValues", func(t *testing.T) {
		testutils.SendChan(channel, "Hello world! Again!")
		testutils.RequireChanEqual(t, channel, "Hello world! Again!")
	})
}
