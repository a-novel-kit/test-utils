package testutils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// CaptureChan captures the output of a channel and stores it in a pointer.
func CaptureChan[T any](channel <-chan T, dest *T) func() {
	interrupt := make(chan bool)

	// Read initial value.
	select {
	case *dest = <-channel:
	default:
	}

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

func RequireChanEqual[T comparable](tb testing.TB, channel <-chan T, expected T) {
	tb.Helper()

	var dest T
	clean := CaptureChan(channel, &dest)
	defer clean()

	require.Eventuallyf(tb, func() bool {
		return dest == expected
	}, time.Second, 50*time.Millisecond, "Expected %q, got %q", expected, dest)
}
