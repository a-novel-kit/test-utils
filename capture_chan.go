package testutils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// CaptureChan captures the output of a channel and stores it in a pointer.
func CaptureChan[T any](channel <-chan T, dest *T) func() {
	interrupt := make(chan bool)

	// Read initial value, if any.
	*dest = <-channel

	go func() {
		for {
			select {
			case *dest = <-channel:
			case <-interrupt:
				return
			}
		}
	}()

	return func() {
		interrupt <- true
	}
}

// SendChan sends a value to a channel, in a separate goroutine, to prevent blocking.
func SendChan[T any](channel chan<- T, value T) {
	go func() {
		channel <- value
	}()
}

func RequireChanC[T any](
	tb testing.TB,
	channel <-chan T,
	condition func(collect *assert.CollectT, value T),
	timeout time.Duration,
	interval time.Duration,
) {
	tb.Helper()

	var dest T
	clean := CaptureChan(channel, &dest)
	defer clean()

	require.EventuallyWithT(tb, func(collect *assert.CollectT) {
		condition(collect, dest)
	}, timeout, interval)
}

func RequireChan[T any](tb testing.TB, channel <-chan T, condition func(collect *assert.CollectT, value T)) {
	tb.Helper()

	RequireChanC(tb, channel, condition, time.Second, 50*time.Millisecond)
}
